package values

import "fmt"

//Period is class of time scope
type Period struct {
	start Date
	end   Date
}

//NewPeriod function creates new Period instance
func NewPeriod(from, to Date) Period {
	if from.IsZero() || to.IsZero() {
		return Period{start: from, end: to}
	}
	if from.After(to) {
		return Period{start: to, end: from}
	}
	return Period{start: from, end: to}
}

//Equal method returns true if all elemnts in Period instance equal
func (lp Period) Equal(rp Period) bool {
	return lp.start == rp.start && lp.end == rp.end
}

//Contains method returns true if scape of this contains date of parameter.
func (p Period) Contains(dt Date) bool {
	if dt.IsZero() {
		return false
	}
	if p.start.IsZero() && p.end.IsZero() {
		return true
	}
	if p.end.IsZero() {
		return !p.start.After(dt)
	}
	if p.start.IsZero() {
		return !p.end.Before(dt)
	}
	return !p.start.After(dt) && !p.end.Before(dt)
}

//StringStart method returns string of Period.start
func (p Period) StringStart() string {
	if p.start.IsZero() {
		return ""
	}
	return p.start.String()
}

//StringEnd method returns string of Period.end
func (p Period) StringEnd() string {
	if p.end.IsZero() {
		return ""
	}
	return p.end.String()
}

//String method is fmt.Stringer for Period
func (p Period) String() string {
	return fmt.Sprintf("%s - %s", p.StringStart(), p.StringEnd())
}

/* Copyright 2020 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
