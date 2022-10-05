package main

// Generates a list of regular expressions to block based on the settings in the config object.
func BuildBlacklist(cfg *Config) []string {
	var list []string

	/*
		cdpsvc.lgtvcommon.com
		rdx2.lgtvsdp.com
		smartshare.lgtvsdp.com
		sp.pluto.tv.
		tags.tiqcdn.com.
	*/

	// lgtvonline.lge.com cannot be blocked, it is used for internet connection check

	if cfg.Blocking.SmartAd {
		list = append(list, "\\bad\\.lgsmartad\\.com\\b")   // region.ad.lgsmartad.com
		list = append(list, "\\binfo\\.lgsmartad\\.com\\b") // region.info.lgsmartad.com
	}

	if cfg.Blocking.HomeDashboard {
		list = append(list, "\\b(?i)(eic)\\.recommend\\.lgtvcommon\\.com\\b") // eic.recommend.lgtvcommon.com
		list = append(list, "\\b(?i)(eic)\\.service\\.lgtvcommon\\.com\\b")   // eic.service.lgtvcommon.com
		list = append(list, "\\b(?i)(eic)\\.rdl\\.lgtvcommon\\.com\\b")       // eic.rdl.lgtvcommon.com
		list = append(list, "\\b(?i)(eic)\\.homeprv\\.lgtvcommon\\.com\\b")   // eic.homeprv.lgtvcommon.com
		list = append(list, "\\b(?i)(eic)\\.nudge\\.lgtvcommon\\.com\\b")     // eic.nudge.lgtvcommon.com

		list = append(list, "\\blgtvsdp\\.com\\b") // region.lgtvsdp.com

		list = append(list, "\\bibsstat\\.lgappstv\\.com\\b") // region.ibsstat.lgappstv.com

		list = append(list, "\\bwww\\.ueiwsp\\.com\\b") // www.ueiwsp.com
	}

	if cfg.Blocking.Sports {
		// Blocks the "Sports Alert" bar in the Home Dashboard
		list = append(list, "\\b(?i)(eic)\\.sports\\.lgtviot\\.com\\b") // eic.sports.lgtviot.com
	}

	if cfg.Blocking.AppStore {
		list = append(list, "\\blgeapi\\.com\\b")                  // region.lgeapi.com
		list = append(list, "\\blgrecommends\\.lgappstv\\.com\\b") // region.lgrecommends.lgappstv.com
		// Thumbnails in the LG Content Store
		list = append(list, "\\bngfts\\.lge\\.com\\b") // ngfts.lge.com
	}

	if cfg.Blocking.InternetChannels {
		list = append(list, "\\b(?i)(eic)\\.cdpbeacon\\.lgtvcommon\\.com\\b")   // eic.cdpbeacon.lgtvcommon.com
		list = append(list, "\\b(?i)(eic)\\.cdplauncher\\.lgtvcommon\\.com\\b") // eic.cdplauncher.lgtvcommon.com
	}

	if cfg.Blocking.LGIOT {
		list = append(list, "\\b(?i)(eic)\\.lgtviot\\.com\\b")        // eic.lgtviot.com
		list = append(list, "\\b(?i)(eic)\\.api\\.lgtviot\\.com\\b")  // eic.api.lgtviot.com
		list = append(list, "\\b(?i)(eic)\\.push\\.lgtviot\\.com\\b") // eic.push.lgtviot.com
	}

	if cfg.Blocking.Amazon {
		list = append(list, "\\bamazon\\.com\\b")            // *.amazon.com
		list = append(list, "\\bamazonvideo\\.com\\b")       // *.amazonvideo.com
		list = append(list, "\\bmedia-amazon\\.com\\b")      // *.media-amazon.com
		list = append(list, "\\bssl-images-amazon\\.com\\b") // *.ssl-images-amazon.com
		// We could also block *.amazonaws.com for the full un-amazoned experience,
		// but that would break a lot of other things
	}

	if cfg.Blocking.PhilipsHue {
		// Philips Hue's SSDP discovery protocol uses the following domain:
		list = append(list, "\\bdiscovery\\.meethue\\.com\\b") // discovery.meethue.com
	}

	if cfg.Blocking.SoftwareUpdates {
		// Firmware update check
		list = append(list, "\\bsnu\\.lge\\.com\\b") // snu.lge.com
		// Firmware update download
		list = append(list, "\\bsu-ssl\\.lge\\.com\\b") // su-ssl.lge.com
	}

	list = append(list, "\\bemp\\.lgsmartplatform\\.com\\b") // region.emp.lgsmartplatform.com

	return list
}
