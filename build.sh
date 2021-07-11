#!/usr/bin/env bash

package_split=(${package//\// })
package_name="pingbot"
platforms=(
  "linux/amd64"
)

for platform in "${platforms[@]}"
do
  platform_split=(${platform//\// })
  GOOS=${platform_split[0]}
  GOARCH=${platform_split[1]}

  output_name=$package_name'-'$GOOS'-'$GOARCH

  if [ $GOOS = "windows" ]; then
      output_name+='.exe'
  fi

  echo "Building on ${platform}..."

  GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name

  echo "Builded!"
done

md5sum $package_name-* > MD5SUM
sha256sum $package_name-* > SHA256SUM

rm -rf VERSION version.go

cat > version.go << EOF
package main

import (
	"fmt"

	"gitlab.com/gaming0skar123/go/pingbot/config"
)

func main() {
	fmt.Print(config.Version)
}
EOF

go run version.go >> VERSION

rm -rf version.go
