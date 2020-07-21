/*
   Copyright 2020 SUSE

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package mits

// TestsConfig represents the root of the tests configuration document.
type TestsConfig struct {
	MariaDB    TestConfig `yaml:"mariadb"`
	MySQL      TestConfig `yaml:"mysql"`
	PostgreSQL TestConfig `yaml:"postgresql"`
	Redis      TestConfig `yaml:"redis"`
}

// TestConfig represents the configuration for an individual test.
type TestConfig struct {
	Enabled bool `yaml:"enabled"`
}