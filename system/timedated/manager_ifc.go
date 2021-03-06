/*
 * Copyright (C) 2017 ~ 2018 Deepin Technology Co., Ltd.
 *
 * Author:     jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package timedated

import (
	"pkg.deepin.io/lib/dbus1"
	"pkg.deepin.io/lib/dbusutil"
)

// SetTime set the current time and date,
// pass a value of microseconds since 1 Jan 1970 UTC
func (m *Manager) SetTime(sender dbus.Sender, usec int64, relative bool, msg string) *dbus.Error {
	err := m.checkAuthorization("SetTime", msg, sender)
	if err != nil {
		return dbusutil.ToError(err)
	}

	// TODO: check usec validity
	err = m.core.SetTime(usec, relative, false)
	return dbusutil.ToError(err)
}

// SetTimezone set the system time zone, the value from /usr/share/zoneinfo/zone.tab
func (m *Manager) SetTimezone(sender dbus.Sender, timezone, msg string) *dbus.Error {
	err := m.checkAuthorization("SetTimezone", msg, sender)
	if err != nil {
		return dbusutil.ToError(err)
	}

	// TODO: check timezone validity
	if m.core.Timezone.Get() == timezone {
		return nil
	}
	err = m.core.SetTimezone(timezone, false)
	return dbusutil.ToError(err)
}

// SetLocalRTC to control whether the RTC is the local time or UTC.
func (m *Manager) SetLocalRTC(sender dbus.Sender, enabled bool, fixSystem bool, msg string) *dbus.Error {
	err := m.checkAuthorization("SetLocalRTC", msg, sender)
	if err != nil {
		return dbusutil.ToError(err)
	}

	if m.core.LocalRTC.Get() == enabled {
		return nil
	}
	err = m.core.SetLocalRTC(enabled, fixSystem, false)
	return dbusutil.ToError(err)
}

// SetNTP to control whether the system clock is synchronized with the network
func (m *Manager) SetNTP(sender dbus.Sender, enabled bool, msg string) *dbus.Error {
	err := m.checkAuthorization("SetNTP", msg, sender)
	if err != nil {
		return dbusutil.ToError(err)
	}

	if m.core.NTP.Get() == enabled {
		return nil
	}
	err = m.core.SetNTP(enabled, false)
	return dbusutil.ToError(err)
}
