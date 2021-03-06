## [Unreleased]

[3.2.17] 2018-05-29
*   fix(session-daemon): some data race problems
*   feat(appearance): sync desktop backgrounds during startup
*   fix(dock): panic if winInfo.wmClass is nil
*   chore(dock): entry.attachWindow print window info
*   chore: update makefile
*   chore: update makefile for arch `sw_64`
*   fix(gesture): disabled if session inactive

[3.2.16] 2018-05-24
*   add fprintd depends in `Desktop edition system`
*   fix(network): allow to delete when creating vpn connection

[3.2.15] 2018-05-15
*   chore(debian): update build-depends

[3.2.14] 2018-05-14
*   feat(apps): record the launched state of the removed app
*   auto sync po files from transifex
*   feat(appearances): set standard font as monospace font fallback
*   fix(appearance): cursor size of window border is small
*   chore(housekeeping): use go-dbus-factory
*   fix(bluetooth): remove adapters and devices config
*   chore(launcher): move launcher module to dde-session-daemon
*   fix(bluetooth): adapter powered not saved
*   auto sync po files from transifex
*   refactor(bluetooth): refactor code again
*   refactor(bluetooth): refactor code
*   feat(bluetooth): add signal Cancelled
*   chore(bluetooth): use go-dbus-factory
*   chore(appearance): use go-dbus-factory
*   chore(audio): use go-dbus-factory
*   chore(fprintd) use go-dbus-factory
*   chore(systeminfo): use go-dbus-factory
*   chore(timedate): use go-dbus-factory
*   chore(gesture): use go-dbus-factory
*   chore(screenedge): use go-dbus-factory
*   chore(keybinding): use go-dbus-factory
*   fix(apps): directory permissions is not 0755
*   chore(sessionwatcher): use go-dbus-factory
*   chore(session/power): use go-dbus-factory
*   feat: add UI unified authentication service
*   fix(session/power): submodule name typo
*   fix(session/power): submodule name typo
*   fix(network): close hotspot no send notification
*   feat(default-terminal): remove --launch-app option
*   feat(network): ConnectionSession add method SetKeyFd
*   feat(keybinding): allow volume to be adjusted to maximum 150%
*   feat: add apps.com.wechat.web to window_patterns
*   feat(appearance): limit the number of custom wallpapers
*   fix(miracast): failed to emit signal Added and Removed

[3.2.13] 2018-03-28
*   chore(dock): add window pattern for gdevelop
*   fix(appearance): add rgba seetings for wine

[3.2.12] 2018-03-22
*   auto sync po files from transifex
*   feat(dock): add window identify for org.deepin.flatdeb.*
*   refactor: improve english
*   refactor(miracast): use newly lib dbusutil
*   fix(session-daemon): different modules startup sequence

[3.2.11] 2018-03-19
*   auto sync po files from transifex
*   fix(audio): nil pointer error in handleCardEvent
*   refactor(session-daemon): use newly lib dbusutil
*   refactor(bluetooth): use newly lib dbusutil
*   fix(accounts): get blurred image without compare change time
*   refactor(fprintd): use newly lib dbusutil
*   refactor(audio): use newly lib dbusutil
*   refactor(inputdevices): use newly lib dbusutil
*   refactor(appearance): use newly lib dbusutil
*   fix(network): allow to delete when creating connection
*   fix(network): fix device mac address unchanged after set it to empty
*   refactor(keybinding): use newly lib dbusutil
*   fix(network): filter notify if device disabled
*   refactor(mime): use newly lib dbusutil
*   refactor(timedate): use newly lib dbusutil
*   refactor(screenedge): use newly lib dbusutil
*   refactor(sessionwatcher): use newly lib dbusutil
*   refactor(systeminfo): use newly lib dbusutil
*   refactor(screensaver): use newly lib dbusutil
*   refactor(session/power): use lib dbusutil
*   chore: use lib dbusutil new api

[3.2.10] 2018-03-07
*   auto sync po files from transifex
*   refactor(dock): optimize design
*   fix(accounts): replace plaintext with ciphertext when set passwd
*   fix(system-daemon): missing the method ScalePlymouth
*   chore: only enable systemd service
*   fix(lockservice): fix event crash after the frequent unlocking
*   feat(session-init): use newly lib dbusutil
*   refactor: remove dbusutil.PropsMaster
*   feat(network): add l2tp ipsec ike/esp settings
*   Revert "feat(session/power): set dpms off before suspend"
*   auto sync po files from transifex
*   fix(network): fix add connection failed if no activated
*   fix(network): correct wired ip unavailable notification
*   feat: make calltrace as module
*   feat(system-daemon): use newly lib dbusutil
*   fix(default-terminal): can not handle the -e option
*   feat(langselector): replace PropsMu with PropsMaster
*   feat(grub2): replace PropsMu with PropsMaster
*   fix(timedate): fix polkit message untranslated
*   fix: optimize channel statements
*   feat(swapsched): add blkio controller
*   feat(dock): window flash supported
*   refactor(debug): watch cpu/mem anormaly
*   fix(soundeffect): property name Enabled typo
*   feat(soundeffect): use newly lib dbusutil
*   feat(search): use newly lib dbusutil
*   feat(langselector): use newly lib dbusutil
*   feat(grub2): use newly lib dbusutil
*   feat(dde-lockservice): use newly lib dbusutil
*   feat(dde-greeter-setter): use newly lib dbusutil
*   feat(`backlight_helper`): use newly lib dbusutil
*   feat: add calltrace to dump runtime stack
*   chore(translations): update translation source
*   chore(accounts): correct policy translations
*   chore: correct network translations
*   chore: update license
*   chore: add accounts systemd service file
*   chore: move bluez and fprintd to optional dependencies
*   feat(trayicon): merge damage notify events
*   fix(session/power): method StartupNotify appears in the DBus interface
*   fix(accounts): change user config path
*   feat: use new lib gsettings
*   feat(keybinding): regrabAll only after keyboard layout changed
*   fix(dock): dock not show if launcher shown
*   fix: optimize appearance gsettings signal
*   refactor(accounts): elaborate login related action
*   feat(accounts): improve user auth action
*   fix: terminal opened by dde-file-manager work dir is wrong
*   feat: use tool deepin-policy-ts-convert to handle the
*   docs: `add service_trigger.md`
*   feat: dde-session-daemon add new module `service_trigger`

## [3.2.9] - 2018-01-24
*   inputdevices: use imwheel to speed up scrolling
*   langselector: use new lib `language_support`
*   dstore: fix waitJobDone for install job
*   swapsched: fix exec cgdelete error
*   keybinding: eliminate keystroke conflict during startup
*   fix: Adapt lintian
*   inputdevices: fix typo in write imwheel config file
*   network: fix nm code generate failure
*   network: add wifi security type 'wpa-eap'
*   inputdevices: fix property WheelSpeed is not writeable
*   network: optimize the method of updating active connections
*   accounts: add DesktopBackgrounds property for user
*   network: use lib notify
*   swapsched: fix missing service file
*   grub2: no json config file
*   accounts: do not verify desktop background file
*   keybinding: run cmd begin with dbus-send directly
*   session/power: remove too much debug print
*   swapsched: create cgroup sessionID@dde/DE
*   dde-session-init: add module `x_event_monitor`
*   lockservice: auto quit to release resources
*   lockservice: fix access m.authUserTable without lock
*   auto sync po files from transifex
*   network: add new empty functions for NM 1.10.2
*   keybinding: update wm switch interface
*   keybinding: update `system_actions.json`
*   logind: fix json marshal failed in shenwei
*   appearance: fix font filter wrong in arm
*   swapsched: use lib cgroup
*   grub2: fix always call generateThemeBackground
*   session/power: adjust brightness function can be controlled by gsettings
*   modify ldfflags args, fix debug version not work
*   grub2: fix typo error
*   fix compile failed using gccgo
*   keybinding: fix ShortcutManager.keyKeystrokeMap concurrent read and write
*   appearance: delete background also delete blurred
*   accounts: generate new blur image if source file new then blurred
*   auto sync po files from transifex
*   launch default terminal via desktop

## [3.2.8] - 2017-12-13
*   add moudle swapsched
*   doc: update bluetooth faq
*   audio: fix update props after config applied
*   dock: fix method RequestDock ignore param index
*   launcher: add methods GetDisableScaling and SetDisableScaling
*   audio: filter out sound effect sink input
*   launcher: fix can not search for newly installed apps
*   appearance: support java scale
*   appearance: fix pam environment settings be override
*   support networkmanager 1.10
*   session/power: set dpms off before suspend
*   makefile GOLDFLAGS remove libcanberra, debian/control depends remove libcanberra-dev

## [3.2.7] - 2017-11-28
*   gesture: check keyboard grab status before do action
*   mime: add multi default app id
*   audio: select best port if config non-exist
*   plymouth: support ssd theme checker
*   dock: fix index in signal EntryAdded is wrong


## [3.2.6] - 2017-11-16
*   add flatpak to recommends

## [3.2.5] - 2017-11-16
*   audio: remove style in font config
*   network: fix wireless disconnect when delete inactive hotspot
*   logined: update 'UserList' when session removed
*   network: remove autoconnect from wireless hotspot
*   appearance: fix fonts memory used large when loading
*   audio: add switcher to decide whether auto switch port


## [3.2.4] - 2017-11-09
#### Features
*   add com.deepin.daemon.ImageBlur interface

#### Bug Fixes
*   not show newly installed wechat in launcher
*   failed to set some bmp image file as icon
*   the Accels field of two shortcuts is empty

#### Changed
*   make `install_to_hicolor.py` compatibility with older python3


## [3.2.3] - 2017-11-03
#### Features
*   automatic switch port when card changed
*   add shortcut for deepin-system-monitor and color-picker
*   support deepin qt theme settings
*   add touchpad tap gesture
*   add flatpak app window identify method


#### Bug Fixes
*   fix gccgo compile failed
*   fix syndaemon pid file not created
*   fix wireless not work after multiple toggle hotspot
*   fix active connections not updated when deleted the last connection
*   update font config xml version


#### Changed
*   refactor grub theme dbus interface
*   rename 'Logout' shortcut to 'Shutdown Interface'
*   add dependency 'dnsmasq'
*   update notifications for scale setting


##  [3.2.2] - 2017-10-27
#### Features
*   keybinding:  process grab pointer failed ([328aa07a](328aa07a))
*   add fprintd module ([1469e2d4](1469e2d4))

#### Bug Fixes
*   fix fprint dependencies missing ([22dc0735](22dc0735))
*   langselector:  write the configuration file wrong ([ee018ea2](ee018ea2))

#### Changed
*   network: remove band settings from hotspot
*   add proxychains-ng as suggested dependency


## [3.2.1] - 2017-10-25
#### Bug Fixes
*   launcher: RequestUninstall does not remove desktop file in autostart directory ([24d1b698](24d1b698))
*   grub2 policykit message not using user's locale ([aa461794](aa461794))
*   keybinding: failed to handle GSettings changed event correctly ([7583b35b](7583b35b))
*   network: delete dot at end ([800eb0c4](800eb0c4))
*   appearance: Fix scale set failed if file not found ([61b72897](61b72897))
*   keybinding can not use key Delete to delete keystroke ([deae5285](deae5285))

#### Features
*   support setting plymouth scale ([842a080e](842a080e))
*   add fprintd module ([1469e2d4](1469e2d4))
*   keybinding: AddCustomShortcut returns id and type of newly created shortcut ([d74f34f8](d74f34f8))
*   accounts: Add no password login ([b87c7448](b87c7448))
*   keybinding: update screenshot command ([64f62269](64f62269))
*   appearance: theme thumbnail support display scaling ([7cba49d6](7cba49d6))
*   dock: menu of entry add item "Force Quit" ([7b853187](7b853187))
*   appearance: Update greeter config when setting scale ([f1b37a80](f1b37a80))
*   network: Implement routes methods ([6889c2d3](6889c2d3))
*   Add 'dde-greeter-setter' ([4dd38e68](4dd38e68))

#### Changed
*   iw: replace 'iw' command with libnl


## [3.2.0] - 2017-10-12
#### Features
* Add scale factor setter
* Add touchpad palm setter
* Add 'Timedated' module to reduce authorization times
* Add the timer of detecting filesystem left space
* Add the methods of managing proxychains proxy
* Add the method of refreshing wireless list
* Add 'ClonedAddress' property to indicate current network device mac address

#### Changed
* Replace 'xfce/clipboard' with 'gnome/clipboard'
* Refactor 'keybinding' module, replace 'xgb' with 'go-x11-client'
* Update network event notify messages
* Update license
* Reset gesture event state when recieved the end event
* Support to hide apps by modify gsettings
* Support to uninstall 'deepin-fpapp-*' package
* Set the default font style when changing font
* Adjust network widgets layout

#### Bug Fixes
* Fix the bug of detecting network device properties error
* Fix activate network hotspot failed
* Fix 'SetProxy' failed if port is empty
