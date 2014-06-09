//handle LidSwitch, PowerButton and Battery status event.
package power

import "os/exec"
import "dbus/com/deepin/sessionmanager"
import "dbus/org/freedesktop/upower"
import "dbus/org/freedesktop/login1"
import "time"
import "dbus/com/deepin/daemon/keybinding"
import "dbus/com/deepin/daemon/display"
import "fmt"

const (
	//sync with com.deepin.daemon.power.schemas
	ActionBlank       int32 = 0
	ActionSuspend           = 1
	ActionShutdown          = 2
	ActionHibernate         = 3
	ActionInteractive       = 4
	ActionNothing           = 5
	ActionLogout            = 6
)

func doLock() {
	if m, err := sessionmanager.NewSessionManager("com.deepin.SessionManager", "/com/deepin/SessionManager"); err != nil {
		LOGGER.Warning("can't build SessionManager Object:", err)
	} else {
		if err = m.RequestLock(); err != nil {
			LOGGER.Warning("Lock failed:", err)
		}
		sessionmanager.DestroySessionManager(m)
	}

}

func doShowLowpower() {
	exec.Command("/usr/lib/deepin-daemon/dde-lowpower").Start()
}
func doCloseLowpower() {
	exec.Command("killall", "dde-lowpower").Start()
}

func doShutDown() {
	if m, err := sessionmanager.NewSessionManager("com.deepin.SessionManager", "/com/deepin/SessionManager"); err != nil {
		LOGGER.Warning("can't build SessionManager Object:", err)
	} else {
		if err = m.RequestShutdown(); err != nil {
			LOGGER.Warning("Shutdown failed:", err)
		}
		sessionmanager.DestroySessionManager(m)
	}
}

func doSuspend() {
	if m, err := sessionmanager.NewSessionManager("com.deepin.SessionManager", "/com/deepin/SessionManager"); err != nil {
		LOGGER.Warning("can't build SessionManager Object:", err)
	} else {
		if err = m.RequestSuspend(); err != nil {
			LOGGER.Warning("Suspend failed:", err)
		}
		LOGGER.Info("RequestSuspend...", err)
		sessionmanager.DestroySessionManager(m)
	}
}

func doLogout() {
	if m, err := sessionmanager.NewSessionManager("com.deepin.SessionManager", "/com/deepin/SessionManager"); err != nil {
		LOGGER.Warning("can't build SessionManager Object:", err)
	} else {
		if err = m.Logout(); err != nil {
			LOGGER.Warning("ShutDown failed:", err)
		}
		sessionmanager.DestroySessionManager(m)
	}
}

func doShutDownInteractive() {
	exec.Command("dde-shutdown").Start()
}

func (up *Power) handlePowerButton() {
	switch up.PowerButtonAction.Get() {
	case ActionInteractive:
		doShutDownInteractive()
	case ActionShutdown:
		doShutDown()
	case ActionSuspend:
		doSuspend()
	case ActionNothing:
	default:
		LOGGER.Warning("invalid LidSwitchAction:", up.LidClosedAction)
	}
}

func (up *Power) handleLidSwitch(opend bool) {
	if opend {
		fmt.Println("LidOpend...")
		//TODO: DPMS ON
	} else {
		fmt.Println("LidClosed...")
		//TODO: DPMS OFF
		switch up.LidClosedAction.Get() {
		case ActionInteractive:
			doShutDownInteractive()
		case ActionSuspend:
			if isMultihead() && !up.coreSettings.GetBoolean("lid-close-suspend-with-external-monitor") {
				LOGGER.Info("Prevent suspend when lidclosed because another monitor connected")
				return
			}
			doSuspend()
		case ActionShutdown:
			doShutDown()
		case ActionNothing:
		default:
			LOGGER.Warning("invalid LidSwitchAction:", up.LidClosedAction.Get())
		}
	}
}

func isMultihead() bool {
	if dp, err := display.NewDisplay("com.deepin.daemon.Display", "/com/deepin/daemon/Display"); err != nil {
		LOGGER.Error("Can't build com.deepin.daemon.Display Object:", err)
		return false
	} else {
		paths := dp.Monitors.Get()
		if len(paths) > 1 {
			return true
		} else if len(paths) == 1 {
			if m, err := display.NewMonitor("com.deepin.daemon.Display", paths[0]); err != nil {
				return false
			} else if m.IsComposited.Get() {
				return true
			} else {
				return false
			}
		}
	}
	return false
}

func (p *Power) initEventHandle() {
	up, err := upower.NewUpower(UPOWER_BUS_NAME, "/org/freedesktop/UPower")
	if err != nil {
		LOGGER.Error("Can't build org.freedesktop.UPower:", err)
	} else {
		up.ConnectChanged(func() {
			currentLidClosed := up.LidIsClosed.Get()
			if p.lidIsClosed != currentLidClosed {
				p.lidIsClosed = currentLidClosed
				p.handleLidSwitch(!currentLidClosed)
			}
			p.lidIsClosed = currentLidClosed

		})
	}

	mediaKey, err := keybinding.NewMediaKey("com.deepin.daemon.KeyBinding", "/com/deepin/daemon/MediaKey")
	if err != nil {
		LOGGER.Error("Can't build com.deepin.daemon.KeyBinding:", err)
	} else {
		mediaKey.ConnectPowerOff(func(press bool) {
			//prevent mediaKey be destroyed
			mediaKey.DestName = mediaKey.DestName

			if !press {
				p.handlePowerButton()
			}
		})
	}

	login, err := login1.NewManager("org.freedesktop.login1", "/org/freedesktop/login1")
	if err != nil {
		LOGGER.Error("Can't build org.freedesktop.login1:", err)
	} else {
		login.ConnectPrepareForSleep(func(before bool) {
			fmt.Println("Sleep change...", before)
			if before {
				if p.lowBatteryStatus == lowBatteryStatusAction {
					doShowLowpower()
				} else {
					if p.coreSettings.GetBoolean("lock-enabled") {
						doLock()
					}
				}
			} else {
				time.AfterFunc(time.Second*1, func() { p.screensaver.SimulateUserActivity() })
				p.handleBatteryPercentage()
			}
		})
	}
}
