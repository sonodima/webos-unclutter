package main

import (
	"fmt"
	"runtime"

	"github.com/fatih/color"
)

// Prints a colored banner to the console, with information about the runtime.
func PrintHeader() {
	fmt.Print(
		color.HiBlackString(`           `)+color.HiRedString(`  ____   _____`)+"\n",
		color.HiBlackString(`           `)+color.HiRedString(` / __ \ / ____|`)+"\n",
		color.HiBlackString(`  __      _`)+color.HiRedString(`| |  | | (___ `)+color.HiYellowString(` _   _`)+"\n",
		color.HiBlackString(`  \ \ /\ / `)+color.HiRedString(`| |  | |\___ \`)+color.HiYellowString(`| | | |`)+"\n",
		color.HiBlackString(`   \ V  V /`)+color.HiRedString(`| |__| |____) `)+color.HiYellowString(`| |_| |`)+"\n",
		color.HiBlackString(`    \_/\_/  `)+color.HiRedString(`\____/|_____/ `)+color.HiYellowString(`\__,_|`)+"\n\n",
		color.HiBlackString("       OS | "+runtime.GOOS)+"\n",
		color.HiBlackString("     ARCH | "+runtime.GOARCH)+"\n\n",
	)
}
