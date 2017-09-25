
package main

import (
  "toy/lib/orm"
)

func main() {
  orm.CreateTable("students", orm.StringColumn{ "name", 20 })
}
