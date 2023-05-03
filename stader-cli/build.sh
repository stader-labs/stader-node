#!/bin/bash

# This work is licensed and released under GNU GPL v3 or any other later versions. 
# The full text of the license is below/ found at <http://www.gnu.org/licenses/>

# (c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [0.4.0-beta]

# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.

# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.

# You should have received a copy of the GNU General Public License
# along with this program.  If not, see <http://www.gnu.org/licenses/>.

export CGO_ENABLED=0
cd /stader-node/stader-cli

# Build x64 version
GOOS=linux GOARCH=amd64 go build -o stader-cli-linux-amd64 stader-cli.go
GOOS=darwin GOARCH=amd64 go build -o stader-cli-darwin-amd64 stader-cli.go

# Build the arm64 version
GOOS=linux GOARCH=arm64 go build -o stader-cli-linux-arm64 stader-cli.go
GOOS=darwin GOARCH=arm64 go build -o stader-cli-darwin-arm64 stader-cli.go
