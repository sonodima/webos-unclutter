package main

// Contains informations about a domain to block.
type Domain struct {
	// The regular expression to match the domain name.
	Expression string
	// A brief description of the domain.
	Comment string
}

// A list of domains to block.
type BlackList struct {
	Domains []Domain
}

// Adds a domain to the blacklist.
func (b *BlackList) add(domain *Domain) {
	b.Domains = append(b.Domains, *domain)
}

// Creates a new blacklist object from the settings in the configuration.
func NewBlackList(cfg *Config) *BlackList {
	b := &BlackList{}

	// lgtvonline.lge.com cannot be blocked, it is used for internet connection check

	if cfg.Blocking.SmartAd {
		b.add(&Domain{"\\bad\\.lgsmartad\\.com\\b", "LG SmartAd"})        // region.ad.lgsmartad.com
		b.add(&Domain{"\\binfo\\.lgsmartad\\.net\\b", "LG SmartAd Info"}) // region.info.lgsmartad.com
	}

	if cfg.Blocking.HomeDashboard {
		b.add(&Domain{"\\b(?i)(eic)\\.recommend\\.lgtvcommon\\.com\\b", "LG Home Dashboard"}) // eic.recommend.lgtvcommon.com
		b.add(&Domain{"\\b(?i)(eic)\\.rdl\\.lgtvcommon\\.com\\b", "LG Home Dashboard"})       // eic.rdl.lgtvcommon.
		b.add(&Domain{"\\b(?i)(eic)\\.homeprv\\.lgtvcommon\\.com\\b", "LG Home Dashboard"})   // eic.homeprv.lgtvcommon.com
		b.add(&Domain{"\\b(?i)(eic)\\.nudge\\.lgtvcommon\\.com\\b", "LG Home Dashboard"})     // eic.nudge.lgtvcommon.com

		// Called a lot during user contract update
		b.add(&Domain{"\\blgtvsdp\\.com\\b", "LG Home Dashboard"}) // region.lgtvsdp.com

		// Called a lot during user contract update
		b.add(&Domain{"\\b(?i)(eic)\\.service\\.lgtvcommon\\.com\\b", "LG Home Dashboard"}) // eic.service.lgtvcommon.com

		b.add(&Domain{"\\bibsstat\\.lgappstv\\.com\\b", "LG Home Dashboard"}) // region.ibsstat.lgappstv.com

		b.add(&Domain{"\\bwww\\.ueiwsp\\.com\\b", "LG Home Dashboard"}) // www.ueiwsp.com
	}

	if cfg.Blocking.Sports {
		// Blocks the "Sports Alert" bar in the Home Dashboard
		b.add(&Domain{"\\b(?i)(eic)\\.sports\\.lgtviot\\.com\\b", "Sports Alert"}) // eic.sports.lgtviot.com
	}

	if cfg.Blocking.AppStore {
		b.add(&Domain{"\\blgeapi\\.com\\b", "LG AppStore"})                  // region.lgeapi.com
		b.add(&Domain{"\\blgrecommends\\.lgappstv\\.com\\b", "LG AppStore"}) // region.lgrecommends.lgappstv.com
		b.add(&Domain{"\\bngfts\\.lge\\.com\\b", "LG AppStore Thumbnails"})  // ngfts.lge.com
	}

	if cfg.Blocking.InternetChannels {
		b.add(&Domain{"\\b(?i)(eic)\\.cdplauncher\\.lgtvcommon\\.com\\b", "Internet Channels Launcher"}) // eic.cdplauncher.lgtvcommon.com
		b.add(&Domain{"\\b(?i)(eic)\\.cdpbeacon\\.lgtvcommon\\.com\\b", "Internet Channels Beacon"})     // eic.cdpbeacon.lgtvcommon.com
	}

	if cfg.Blocking.LGIOT {
		b.add(&Domain{"\\b(?i)(eic)\\.lgtviot\\.com\\b", "LG IOT"})             // eic.lgtviot.com
		b.add(&Domain{"\\b(?i)(eic)\\.api\\.lgtviot\\.com\\b", "LG IOT API"})   // eic.api.lgtviot.com
		b.add(&Domain{"\\b(?i)(eic)\\.push\\.lgtviot\\.com\\b", "LG IOT Push"}) // eic.push.lgtviot.com
	}

	if cfg.Blocking.Amazon {
		b.add(&Domain{"\\bamazon\\.com\\b", "Amazon"})                   // *.amazon.com
		b.add(&Domain{"\\bamazonvideo\\.com\\b", "Amazon Video"})        // *.amazonvideo.com
		b.add(&Domain{"\\bmedia-amazon\\.com\\b", "Amazon Media"})       // *.media-amazon.com
		b.add(&Domain{"\\bssl-images-amazon\\.com\\b", "Amazon Images"}) // *.ssl-images-amazon.com
		// We could also block *.amazonaws.com for the full un-amazoned experience,
		// but that would break a lot of other things
	}

	if cfg.Blocking.PhilipsHue {
		// Philips Hue's SSDP discovery protocol uses the following domain:
		b.add(&Domain{"\\bdiscovery\\.meethue\\.com\\b", "Philips Hue SSDP"}) // discovery.meethue.com
	}

	if cfg.Blocking.SoftwareUpdates {
		b.add(&Domain{"\\bsnu\\.lge\\.com\\b", "Firmware Update Check"})       // snu.lge.com
		b.add(&Domain{"\\bsu-ssl\\.lge\\.com\\b", "Firmware Update Download"}) // su-ssl.lge.com
	}

	b.add(&Domain{"\\bemp\\.lgsmartplatform\\.com\\b", "LG Smart Platform EMP"}) // region.emp.lgsmartplatform.com

	b.add(&Domain{"\\bcdpsvc\\.lgtvcommon\\.com\\b", "Unknown CDPSVC"}) // cdpsvc.lgtvcommon.com
	b.add(&Domain{"\\brdx2\\.lgtvsdp\\.com\\b", "Unknown RDX2"})        // rdx2.lgtvsdp.com
	b.add(&Domain{"\\bsmartshare\\.lgtvsdp\\.com\\b", "SmartShare"})    // smartshare.lgtvsdp.com
	b.add(&Domain{"\\bsp\\.pluto\\.tv\\b", "Pluto TV"})                 // sp.pluto.tv
	b.add(&Domain{"\\btags\\.tiqcdn\\.com\\b", "Unknown TIQCDN"})       // tags.tiqcdn.com
	b.add(&Domain{"\\bibs\\.lgappstv\\.com\\b", "Unknown IBS"})         // ibs.lgappstv.com

	return b
}
