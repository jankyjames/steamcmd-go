package steamcmd

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testSteamPayload = []byte(`{
	"common"
	{
		"icon"		"e3f595a92552da3d664ad00277fad2107345f743"
		"logo"		"07385eb55b5ba974aebbe74d3c99626bda7920b8"
		"logo_small"		"07385eb55b5ba974aebbe74d3c99626bda7920b8_thumb"
		"metacritic_url"		"pc/teamfortress2"
		"name"		"Team Fortress 2"
		"clienticon"		"033bdd91842b6aca0633ee1e5f3e6b82f2e8962f"
		"clienttga"		"689dd46bd63e3e460cdd86d936d5de409a291633"
		"languages"
		{
			"english"		"1"
			"german"		"1"
			"french"		"1"
			"italian"		"1"
			"koreana"		"1"
			"spanish"		"1"
			"schinese"		"1"
			"tchinese"		"1"
			"russian"		"1"
			"thai"		"1"
			"japanese"		"1"
			"portuguese"		"1"
			"polish"		"1"
			"danish"		"1"
			"dutch"		"1"
			"finnish"		"1"
			"norwegian"		"1"
			"swedish"		"1"
			"hungarian"		"1"
			"czech"		"1"
			"romanian"		"1"
			"turkish"		"1"
			"brazilian"		"1"
			"bulgarian"		"1"
			"greek"		"1"
			"ukrainian"		"1"
		}
		"clienticns"		"b0780e91df4308b1ad57cc93d6032b4f0e2930cc"
		"linuxclienticon"		"b2659c540592221fcd7675d76a0171f4b3782c1c"
		"oslist"		"windows,macos,linux"
		"osarch"		""
		"osextended"		""
		"type"		"game"
		"content_descriptors"
		{
			"0"		"2"
			"1"		"5"
		}
		"steam_deck_compatibility"
		{
			"category"		"2"
			"test_timestamp"		"1713571200"
			"tested_build_id"		"7850740"
			"tests"
			{
				"0"
				{
					"display"		"3"
					"token"		"#SteamDeckVerified_TestResult_DefaultControllerConfigNotFullyFunctional"
				}
				"1"
				{
					"display"		"3"
					"token"		"#SteamDeckVerified_TestResult_ControllerGlyphsDoNotMatchDeckDevice"
				}
				"2"
				{
					"display"		"3"
					"token"		"#SteamDeckVerified_TestResult_InterfaceTextIsNotLegible"
				}
				"3"
				{
					"display"		"4"
					"token"		"#SteamDeckVerified_TestResult_DefaultConfigurationIsPerformant"
				}
				"4"
				{
					"display"		"1"
					"token"		"#SteamDeckVerified_TestResult_ExternalControllersNotSupportedPrimaryPlayer"
				}
			}
			"configuration"
			{
				"supported_input"		"gamepad"
				"requires_manual_keyboard_invoke"		"0"
				"requires_non_controller_launcher_nav"		"0"
				"primary_player_is_controller_slot_0"		"1"
				"non_deck_display_glyphs"		"1"
				"small_text"		"1"
				"requires_internet_for_setup"		"0"
				"requires_internet_for_singleplayer"		"0"
				"recommended_runtime"		"native"
				"requires_h264"		"0"
				"gamescope_frame_limiter_not_supported"		"0"
				"hdr_support"		"0"
			}
		}
		"market_presence"		"1"
		"metacritic_name"		"Team Fortress 2"
		"controller_support"		"partial"
		"small_capsule"
		{
			"english"		"capsule_231x87.jpg"
		}
		"header_image"
		{
			"english"		"header.jpg"
		}
		"library_assets"
		{
			"library_capsule"		"en"
			"library_hero"		"en"
			"library_logo"		"en"
			"logo_position"
			{
				"pinned_position"		"BottomLeft"
				"width_pct"		"26"
				"height_pct"		"37"
			}
		}
		"store_asset_mtime"		"1682961190"
		"associations"
		{
			"0"
			{
				"type"		"developer"
				"name"		"Valve"
			}
			"1"
			{
				"type"		"publisher"
				"name"		"Valve"
			}
		}
		"primary_genre"		"37"
		"genres"
		{
			"0"		"1"
			"1"		"37"
		}
		"category"
		{
			"category_8"		"1"
			"category_12"		"1"
			"category_13"		"1"
			"category_14"		"1"
			"category_15"		"1"
			"category_17"		"1"
			"category_18"		"1"
			"category_22"		"1"
			"category_27"		"1"
			"category_1"		"1"
			"category_29"		"1"
			"category_30"		"1"
			"category_33"		"1"
			"category_35"		"1"
			"category_45"		"1"
			"category_46"		"1"
			"category_41"		"1"
			"category_42"		"1"
		}
		"supported_languages"
		{
			"english"
			{
				"supported"		"true"
				"full_audio"		"true"
			}
			"danish"
			{
				"supported"		"true"
			}
			"dutch"
			{
				"supported"		"true"
			}
			"finnish"
			{
				"supported"		"true"
			}
			"french"
			{
				"supported"		"true"
			}
			"german"
			{
				"supported"		"true"
			}
			"italian"
			{
				"supported"		"true"
			}
			"japanese"
			{
				"supported"		"true"
			}
			"norwegian"
			{
				"supported"		"true"
			}
			"polish"
			{
				"supported"		"true"
			}
			"portuguese"
			{
				"supported"		"true"
			}
			"russian"
			{
				"supported"		"true"
			}
			"schinese"
			{
				"supported"		"true"
			}
			"spanish"
			{
				"supported"		"true"
			}
			"swedish"
			{
				"supported"		"true"
			}
			"tchinese"
			{
				"supported"		"true"
			}
			"koreana"
			{
				"supported"		"true"
			}
			"czech"
			{
				"supported"		"true"
			}
			"hungarian"
			{
				"supported"		"true"
			}
			"brazilian"
			{
				"supported"		"true"
			}
			"turkish"
			{
				"supported"		"true"
			}
			"greek"
			{
				"supported"		"true"
			}
			"bulgarian"
			{
				"supported"		"true"
			}
			"romanian"
			{
				"supported"		"true"
			}
			"thai"
			{
				"supported"		"true"
			}
			"ukrainian"
			{
				"supported"		"true"
			}
		}
		"original_release_date"		"1191999600"
		"steam_release_date"		"1191999600"
		"metacritic_score"		"92"
		"metacritic_fullurl"		"https://www.metacritic.com/game/pc/team-fortress-2?ftag=MCD-06-10aaa1f"
		"community_visible_stats"		"1"
		"workshop_visible"		"1"
		"community_hub_visible"		"1"
		"gameid"		"440"
		"exfgls"		"3"
		"store_tags"
		{
			"0"		"113"
			"1"		"620519"
			"2"		"3859"
			"3"		"1663"
			"4"		"1774"
			"5"		"19"
			"6"		"4155"
			"7"		"5711"
			"8"		"4136"
			"9"		"3839"
			"10"		"3843"
			"11"		"3878"
			"12"		"4195"
			"13"		"4202"
			"14"		"1685"
			"15"		"1719"
			"16"		"5752"
			"17"		"1708"
			"18"		"4562"
			"19"		"1702"
		}
		"review_score"		"8"
		"review_percentage"		"85"
	}
	"extended"
	{
		"developer"		"Valve"
		"developer_url"		"http://www.valvesoftware.com/"
		"gamedir"		"tf"
		"gamemanualurl"		"http://store.steampowered.com/manual/440/"
		"homepage"		"http://www.teamfortress.com/"
		"icon"		"Client/games/icon_tf2"
		"icon2"		"Client/games/icon_tf2"
		"isfreeapp"		"1"
		"languages"		"english,german,french,spanish,russian"
		"loadallbeforelaunch"		"1"
		"minclientversion"		"1393366296"
		"noservers"		"0"
		"primarycache"		"441"
		"primarycache_linux"		"452"
		"requiressse"		"1"
		"serverbrowsername"		"Team Fortress 2"
		"sourcegame"		"1"
		"state"		"eStateAvailable"
		"vacmacmodulecache"		"160"
		"vacmodulecache"		"202"
		"vacmodulefilename"		"sourceinit.dat"
		"validoslist"		"windows,macos,linux"
		"deckresolutionoverride"		"Native"
		"publisher"		"Valve"
		"aliases"		"tf2"
		"listofdlc"		"456,457,458,459,217221,217222"
	}
	"config"
	{
		"verifyupdates"		"1"
		"signedfiles"
		{
			"hl2.exe"		"30819D300D06092A864886F70D010101050003818B0030818702818100B1260881BDFE84463D88C6AB8DB914A2E593893C10508B8A5ABDF692E9A5419A3EDBAE86A052849983B75E3B425C18178B260003D857DF0B6505C6CF9C84F5859FCE3B63F1FB2D4818501F6C5FA4AD1430EEB081A74ABD74CD1F4AA1FCCA3B88DD0548AED34443CEB52444EAE9099AA4FE66B2E6224D02381C248025C7044079020111"
			"Client.dll"		"30819D300D06092A864886F70D010101050003818B0030818702818100B1260881BDFE84463D88C6AB8DB914A2E593893C10508B8A5ABDF692E9A5419A3EDBAE86A052849983B75E3B425C18178B260003D857DF0B6505C6CF9C84F5859FCE3B63F1FB2D4818501F6C5FA4AD1430EEB081A74ABD74CD1F4AA1FCCA3B88DD0548AED34443CEB52444EAE9099AA4FE66B2E6224D02381C248025C7044079020111"
			"engine.dll"		"30819D300D06092A864886F70D010101050003818B0030818702818100B1260881BDFE84463D88C6AB8DB914A2E593893C10508B8A5ABDF692E9A5419A3EDBAE86A052849983B75E3B425C18178B260003D857DF0B6505C6CF9C84F5859FCE3B63F1FB2D4818501F6C5FA4AD1430EEB081A74ABD74CD1F4AA1FCCA3B88DD0548AED34443CEB52444EAE9099AA4FE66B2E6224D02381C248025C7044079020111"
			"server.dll"		"30819D300D06092A864886F70D010101050003818B0030818702818100B1260881BDFE84463D88C6AB8DB914A2E593893C10508B8A5ABDF692E9A5419A3EDBAE86A052849983B75E3B425C18178B260003D857DF0B6505C6CF9C84F5859FCE3B63F1FB2D4818501F6C5FA4AD1430EEB081A74ABD74CD1F4AA1FCCA3B88DD0548AED34443CEB52444EAE9099AA4FE66B2E6224D02381C248025C7044079020111"
			"prec.dll"		"30819D300D06092A864886F70D010101050003818B0030818702818100BA240326DFC66BFDC28782BD644FC3BAB284F4A07554E17A8DA6CCA0B4AFBFC9BAB2938846D7595BC7A560FF6215C670C74EEB9D118DE9FC0E8ED0A18DD24D8DBAE8B950C6C062BA5DEFDECB967745C53433839EB7B1A6CA340D0E48273E18AB87106FBD22AC32B2776092E924D3C414DC3328735E984E4A1B1B4E72B0F41017020111"
			"sourcevr.dll"		"30819D300D06092A864886F70D010101050003818B0030818702818100B1260881BDFE84463D88C6AB8DB914A2E593893C10508B8A5ABDF692E9A5419A3EDBAE86A052849983B75E3B425C18178B260003D857DF0B6505C6CF9C84F5859FCE3B63F1FB2D4818501F6C5FA4AD1430EEB081A74ABD74CD1F4AA1FCCA3B88DD0548AED34443CEB52444EAE9099AA4FE66B2E6224D02381C248025C7044079020111"
		}
		"usemms"		""
		"systemprofile"		"1"
		"checkforupdatesbeforelaunch"		"1"
		"installdir"		"Team Fortress 2"
		"launch"
		{
			"0"
			{
				"executable"		"tf_win64.exe"
				"arguments"		"-Client"
				"config"
				{
					"oslist"		"windows"
					"osarch"		"64"
				}
			}
			"1"
			{
				"executable"		"hl2_osx"
				"arguments"		"-Client -game tf"
				"config"
				{
					"oslist"		"macos"
				}
			}
			"2"
			{
				"executable"		"tf.sh"
				"arguments"		"-Client"
				"config"
				{
					"oslist"		"linux"
					"osarch"		"64"
				}
			}
			"6"
			{
				"executable"		"tf.exe"
				"arguments"		"-Client"
				"config"
				{
					"oslist"		"windows"
					"osarch"		"32"
				}
			}
			"7"
			{
				"executable"		"tf.sh"
				"arguments"		"-Client -gl"
				"config"
				{
					"oslist"		"linux"
					"osarch"		"64"
				}
				"description_loc"
				{
					"english"		"Legacy OpenGL"
				}
				"description"		"Legacy OpenGL"
			}
		}
		"vrcompositorsupport"		"0"
		"steamcontrollertemplateindex"		"1"
		"steamcontrollerconfigdetails"
		{
			"1172518660"
			{
				"controller_type"		"controller_steamcontroller_gordon"
				"enabled_branches"		"default"
			}
		}
		"steamcontrollertouchtemplateindex"		"1"
		"steamcontrollertouchconfigdetails"
		{
			"2125279140"
			{
				"controller_type"		"controller_mobile_touch"
				"enabled_branches"		"default"
				"use_action_block"		"false"
			}
		}
		"uselaunchcommandline"		"1"
		"app_mappings"
		{
			"1"
			{
				"platform"		"linux"
				"tool"		"SteamLinuxRuntime_sniper"
				"comment"		"sniper SLR"
				"priority"		"101"
			}
		}
	}
	"depots"
	{
		"overridescddb"		"1"
		"baselanguages"		"english,german,french,italian,koreana,spanish,schinese,tchinese,russian,thai,japanese,portuguese,polish,danish,dutch,finnish,norwegian,swedish,hungarian,czech,romanian,turkish,brazilian,bulgarian,greek,ukrainian"
		"workshopdepot"		"440"
		"hasdepotsindlc"		"1"
		"228990"
		{
			"config"
			{
				"oslist"		"windows"
			}
			"depotfromapp"		"228980"
			"sharedinstall"		"1"
		}
		"441"
		{
			"systemdefined"		"1"
			"manifests"
			{
				"public"
				{
					"gid"		"8108861274672275806"
					"size"		"27585185992"
					"download"		"13104915840"
				}
				"pre_jungleinferno_demos"
				{
					"gid"		"7707612755534478338"
					"size"		"18855865767"
					"download"		"7735778976"
				}
				"prerelease"
				{
					"gid"		"8108861274672275806"
					"size"		"27585185992"
					"download"		"13104915840"
				}
				"pre_07_25_23_demos"
				{
					"gid"		"2023276112389133529"
					"size"		"26125661395"
					"download"		"11978937568"
				}
				"pre_smissmas_2022_demos"
				{
					"gid"		"314077012305764573"
					"size"		"24749544116"
					"download"		"10985537760"
				}
			}
		}
		"440"
		{
			"systemdefined"		"1"
			"manifests"
			{
				"public"
				{
					"gid"		"1118032470228587934"
					"size"		"825745"
					"download"		"43168"
				}
				"pre_jungleinferno_demos"
				{
					"gid"		"1118032470228587934"
					"size"		"825745"
					"download"		"43168"
				}
				"prerelease"
				{
					"gid"		"1118032470228587934"
					"size"		"825745"
					"download"		"43168"
				}
				"pre_07_25_23_demos"
				{
					"gid"		"1118032470228587934"
					"size"		"825745"
					"download"		"43168"
				}
				"pre_smissmas_2022_demos"
				{
					"gid"		"1118032470228587934"
					"size"		"825745"
					"download"		"43168"
				}
			}
		}
		"232251"
		{
			"config"
			{
				"oslist"		"windows"
			}
			"systemdefined"		"1"
			"manifests"
			{
				"public"
				{
					"gid"		"5648007322442819081"
					"size"		"553794707"
					"download"		"352973840"
				}
				"pre_jungleinferno_demos"
				{
					"gid"		"2174530283606128348"
					"size"		"562700702"
					"download"		"445092848"
				}
				"prerelease"
				{
					"gid"		"5648007322442819081"
					"size"		"553794707"
					"download"		"352973840"
				}
				"pre_07_25_23_demos"
				{
					"gid"		"1093030316367308660"
					"size"		"585005817"
					"download"		"467050320"
				}
				"pre_smissmas_2022_demos"
				{
					"gid"		"8319568912250137931"
					"size"		"584144460"
					"download"		"466735776"
				}
			}
		}
		"232252"
		{
			"config"
			{
				"oslist"		"macos"
			}
			"systemdefined"		"1"
			"manifests"
			{
				"public"
				{
					"gid"		"678390272316880472"
					"size"		"103246051"
					"download"		"22586512"
				}
				"pre_jungleinferno_demos"
				{
					"gid"		"8624037758464932425"
					"size"		"485750759"
					"download"		"370815712"
				}
				"prerelease"
				{
					"gid"		"678390272316880472"
					"size"		"103246051"
					"download"		"22586512"
				}
				"pre_07_25_23_demos"
				{
					"gid"		"4845664414220472500"
					"size"		"510895405"
					"download"		"391271824"
				}
				"pre_smissmas_2022_demos"
				{
					"gid"		"2262584799279223344"
					"size"		"509176417"
					"download"		"390890000"
				}
			}
		}
		"232253"
		{
			"config"
			{
				"oslist"		"linux"
			}
			"systemdefined"		"1"
			"manifests"
			{
				"public"
				{
					"gid"		"2936685618595780"
					"size"		"405622516"
					"download"		"285119312"
				}
				"pre_jungleinferno_demos"
				{
					"gid"		"3512165917877440762"
					"size"		"536705326"
					"download"		"424011840"
				}
				"prerelease"
				{
					"gid"		"2936685618595780"
					"size"		"405622516"
					"download"		"285119312"
				}
				"pre_07_25_23_demos"
				{
					"gid"		"8481848623592128300"
					"size"		"562481508"
					"download"		"446661232"
				}
				"pre_smissmas_2022_demos"
				{
					"gid"		"2039032268643181297"
					"size"		"560994616"
					"download"		"446299360"
				}
			}
		}
		"444"
		{
			"config"
			{
				"lowviolence"		"1"
			}
			"manifests"
			{
				"public"
				{
					"gid"		"646330807595077556"
					"size"		"7452052"
					"download"		"2986208"
				}
				"pre_jungleinferno_demos"
				{
					"gid"		"646330807595077556"
					"size"		"7452052"
					"download"		"2986208"
				}
				"prerelease"
				{
					"gid"		"646330807595077556"
					"size"		"7452052"
					"download"		"2986208"
				}
				"pre_07_25_23_demos"
				{
					"gid"		"646330807595077556"
					"size"		"7452052"
					"download"		"2986208"
				}
				"pre_smissmas_2022_demos"
				{
					"gid"		"646330807595077556"
					"size"		"7452052"
					"download"		"2986208"
				}
			}
		}
		"445"
		{
			"config"
			{
				"language"		"russian"
			}
			"manifests"
			{
				"public"
				{
					"gid"		"4462275339090266921"
					"size"		"90075889"
					"download"		"79353264"
				}
				"pre_jungleinferno_demos"
				{
					"gid"		"5602233730528152354"
					"size"		"90075889"
					"download"		"79353248"
				}
				"prerelease"
				{
					"gid"		"4462275339090266921"
					"size"		"90075889"
					"download"		"79353264"
				}
				"pre_07_25_23_demos"
				{
					"gid"		"5602233730528152354"
					"size"		"90075889"
					"download"		"79353248"
				}
				"pre_smissmas_2022_demos"
				{
					"gid"		"5602233730528152354"
					"size"		"90075889"
					"download"		"79353248"
				}
			}
		}
		"446"
		{
			"config"
			{
				"language"		"german"
			}
			"manifests"
			{
				"public"
				{
					"gid"		"4932209831621142568"
					"size"		"74594251"
					"download"		"65214112"
				}
				"pre_jungleinferno_demos"
				{
					"gid"		"6215119033293728486"
					"size"		"74594251"
					"download"		"65214128"
				}
				"prerelease"
				{
					"gid"		"4932209831621142568"
					"size"		"74594251"
					"download"		"65214112"
				}
				"pre_07_25_23_demos"
				{
					"gid"		"6215119033293728486"
					"size"		"74594251"
					"download"		"65214128"
				}
				"pre_smissmas_2022_demos"
				{
					"gid"		"6215119033293728486"
					"size"		"74594251"
					"download"		"65214128"
				}
			}
		}
		"448"
		{
			"config"
			{
				"language"		"french"
			}
			"manifests"
			{
				"public"
				{
					"gid"		"638274067337307989"
					"size"		"73934506"
					"download"		"64158176"
				}
				"pre_jungleinferno_demos"
				{
					"gid"		"8469957949853360122"
					"size"		"73934506"
					"download"		"64158192"
				}
				"prerelease"
				{
					"gid"		"638274067337307989"
					"size"		"73934506"
					"download"		"64158176"
				}
				"pre_07_25_23_demos"
				{
					"gid"		"8469957949853360122"
					"size"		"73934506"
					"download"		"64158192"
				}
				"pre_smissmas_2022_demos"
				{
					"gid"		"8469957949853360122"
					"size"		"73934506"
					"download"		"64158192"
				}
			}
		}
		"449"
		{
			"config"
			{
				"language"		"spanish"
			}
			"manifests"
			{
				"public"
				{
					"gid"		"476175847494669680"
					"size"		"73947996"
					"download"		"58909360"
				}
				"pre_jungleinferno_demos"
				{
					"gid"		"2451727451644207550"
					"size"		"73947996"
					"download"		"58909344"
				}
				"prerelease"
				{
					"gid"		"476175847494669680"
					"size"		"73947996"
					"download"		"58909360"
				}
				"pre_07_25_23_demos"
				{
					"gid"		"2451727451644207550"
					"size"		"73947996"
					"download"		"58909344"
				}
				"pre_smissmas_2022_demos"
				{
					"gid"		"2451727451644207550"
					"size"		"73947996"
					"download"		"58909344"
				}
			}
		}
		"branches"
		{
			"public"
			{
				"buildid"		"14136189"
				"timeupdated"		"1713822697"
			}
			"pre_jungleinferno_demos"
			{
				"buildid"		"2119861"
				"timeupdated"		"1690827460"
			}
			"prerelease"
			{
				"buildid"		"14136189"
				"timeupdated"		"1713822753"
			}
			"pre_07_25_23_demos"
			{
				"buildid"		"11753452"
				"description"		"Pre-07-25-23 Demo Playback"
				"timeupdated"		"1690416239"
			}
			"pre_smissmas_2022_demos"
			{
				"buildid"		"9872056"
				"description"		"Pre-Smissmas 2022 Demo Playback"
				"timeupdated"		"1690827525"
			}
		}
	}
	"ufs"
	{
		"quota"		"1073741824"
		"maxnumfiles"		"1000"
		"hidecloudui"		"0"
	}
	"sysreqs"
	{
		"macos"
		{
			"supported"		"1"
			"wants_fast_gpu"		"1"
			"intel900"		"warn"
			"intelx3100"		"warn"
			"8086"
			{
				"27a6"		"warn"
				"2a02"		"warn"
			}
			"10de"
			{
				"009d"		"deny"
				"0391"		"deny"
				"0393"		"deny"
				"0395"		"deny"
			}
		}
		"macos104"
		{
			"supported"		"0"
			"os_min"		"macos1058"
		}
		"macos105"
		{
			"supported"		"0"
			"os_min"		"macos1058"
		}
		"macos106"
		{
			"supported"		"0"
			"os_min"		"macos1063"
		}
	}
	"localization"
	{
		"richpresence"
		{
			"english"
			{
				"tokens"
				{
					"#tf_richpresence_display"		"{#TF_RichPresence_State_%state%}"
					"#tf_richpresence_state_mainmenu"		"Main Menu"
					"#tf_richpresence_state_searchinggeneric"		"Searching for a Match"
					"#tf_richpresence_state_searchingmatchgroup"		"Searching - {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playinggeneric"		"In Match - %currentmap%"
					"#tf_richpresence_state_loadinggeneric"		"Joining Match"
					"#tf_richpresence_state_playingmatchgroup"		"{#TF_RichPresence_MatchGroup_%matchgrouploc%} - %currentmap%"
					"#tf_richpresence_state_loadingmatchgroup"		"Joining {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playingcommunity"		"Community - %currentmap%"
					"#tf_richpresence_state_loadingcommunity"		"Joining Community Server"
					"#tf_richpresence_matchgroup_competitive6v6"		"Competitive"
					"#tf_richpresence_matchgroup_casual"		"Casual"
					"#tf_richpresence_matchgroup_specialevent"		"Special Event"
					"#tf_richpresence_matchgroup_mannup"		"Mann Up"
					"#tf_richpresence_matchgroup_bootcamp"		"Boot Camp"
				}
			}
			"bulgarian"
			{
				"tokens"
				{
					"#tf_richpresence_display"		"{#TF_RichPresence_State_%state%}"
					"#tf_richpresence_state_mainmenu"		"Главно меню"
					"#tf_richpresence_state_searchinggeneric"		"Търсене за мач"
					"#tf_richpresence_state_searchingmatchgroup"		"Търсене — {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playinggeneric"		"В мач — %currentmap%"
					"#tf_richpresence_state_loadinggeneric"		"Присъединяване към мач"
					"#tf_richpresence_state_playingmatchgroup"		"{#TF_RichPresence_MatchGroup_%matchgrouploc%} — %currentmap%"
					"#tf_richpresence_state_loadingmatchgroup"		"Присъединяване — {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playingcommunity"		"Общност — %currentmap%"
					"#tf_richpresence_state_loadingcommunity"		"Присъединяване към обществен сървър"
					"#tf_richpresence_matchgroup_competitive6v6"		"Съревнователна"
					"#tf_richpresence_matchgroup_casual"		"Неангажираща"
					"#tf_richpresence_matchgroup_specialevent"		"Специално събитие"
					"#tf_richpresence_matchgroup_mannup"		"Щателни Манневри"
					"#tf_richpresence_matchgroup_bootcamp"		"Тренировъчен лагер"
				}
			}
			"danish"
			{
				"tokens"
				{
					"#tf_richpresence_display"		"{#TF_RichPresence_State_%state%}"
					"#tf_richpresence_state_mainmenu"		"Hovedmenu"
					"#tf_richpresence_state_searchinggeneric"		"Søger efter en kamp"
					"#tf_richpresence_state_searchingmatchgroup"		"Søger - {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playinggeneric"		"I en kamp - %currentmap%"
					"#tf_richpresence_state_loadinggeneric"		"Slutter til en kamp"
					"#tf_richpresence_state_playingmatchgroup"		"{#TF_RichPresence_MatchGroup_%matchgrouploc%} - %currentmap%"
					"#tf_richpresence_state_loadingmatchgroup"		"Slutter til {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playingcommunity"		"Fællesskab - %currentmap%"
					"#tf_richpresence_state_loadingcommunity"		"Slutter til en fællesskabsserver"
					"#tf_richpresence_matchgroup_competitive6v6"		"Competitive"
					"#tf_richpresence_matchgroup_casual"		"Casual"
					"#tf_richpresence_matchgroup_specialevent"		"Særlig begivenhed"
					"#tf_richpresence_matchgroup_mannup"		"Mann Up"
					"#tf_richpresence_matchgroup_bootcamp"		"Træning"
				}
			}
			"french"
			{
				"tokens"
				{
					"#tf_richpresence_display"		"{#TF_RichPresence_State_%state%}"
					"#tf_richpresence_state_mainmenu"		"Menu principal"
					"#tf_richpresence_state_searchinggeneric"		"Cherche un match"
					"#tf_richpresence_state_searchingmatchgroup"		"Recherche - {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playinggeneric"		"Dans un match - %currentmap%"
					"#tf_richpresence_state_loadinggeneric"		"Rejoint un match"
					"#tf_richpresence_state_playingmatchgroup"		"{#TF_RichPresence_MatchGroup_%matchgrouploc%} - %currentmap%"
					"#tf_richpresence_state_loadingmatchgroup"		"Rejoint un match {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playingcommunity"		"Serveur de la communauté - %currentmap%"
					"#tf_richpresence_state_loadingcommunity"		"Rejoint un serveur de la communauté"
					"#tf_richpresence_matchgroup_competitive6v6"		"Compétitif"
					"#tf_richpresence_matchgroup_casual"		"Occasionnel"
					"#tf_richpresence_matchgroup_specialevent"		"Événement spécial"
					"#tf_richpresence_matchgroup_mannup"		"Mann Up"
					"#tf_richpresence_matchgroup_bootcamp"		"Camp d'entraînement"
				}
			}
			"german"
			{
				"tokens"
				{
					"#tf_richpresence_display"		"{#TF_RichPresence_State_%state%}"
					"#tf_richpresence_state_mainmenu"		"Hauptmenü"
					"#tf_richpresence_state_searchinggeneric"		"Suche nach Spiel läuft"
					"#tf_richpresence_state_searchingmatchgroup"		"Suche läuft – {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playinggeneric"		"Im Spiel – %currentmap%"
					"#tf_richpresence_state_loadinggeneric"		"Spielbeitritt erfolgt"
					"#tf_richpresence_state_playingmatchgroup"		"{#TF_RichPresence_MatchGroup_%matchgrouploc%} – %currentmap%"
					"#tf_richpresence_state_loadingmatchgroup"		"Sie treten {#TF_RichPresence_MatchGroup_%matchgrouploc%} bei"
					"#tf_richpresence_state_playingcommunity"		"Community – %currentmap%"
					"#tf_richpresence_state_loadingcommunity"		"Verbindung zum Communityserver erfolgt"
					"#tf_richpresence_matchgroup_competitive6v6"		"Wettkampf"
					"#tf_richpresence_matchgroup_casual"		"Gelegenheitsspiel"
					"#tf_richpresence_matchgroup_specialevent"		"Sonderevent"
					"#tf_richpresence_matchgroup_mannup"		"Mann-Up"
					"#tf_richpresence_matchgroup_bootcamp"		"Training"
				}
			}
			"hungarian"
			{
				"tokens"
				{
					"#tf_richpresence_display"		"{#TF_RichPresence_State_%state%}"
					"#tf_richpresence_state_mainmenu"		"Főmenüben"
					"#tf_richpresence_state_searchinggeneric"		"Meccset keres"
					"#tf_richpresence_state_searchingmatchgroup"		"Keres - {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playinggeneric"		"Meccsen - %currentmap%"
					"#tf_richpresence_state_loadinggeneric"		"Meccshez csatlakozik"
					"#tf_richpresence_state_playingmatchgroup"		"{#TF_RichPresence_MatchGroup_%matchgrouploc%} - %currentmap%"
					"#tf_richpresence_state_loadingmatchgroup"		"Csatlakozik: {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playingcommunity"		"Közösség - %currentmap%"
					"#tf_richpresence_state_loadingcommunity"		"Közösségi szerverhez csatlakozik"
					"#tf_richpresence_matchgroup_competitive6v6"		"Versengő"
					"#tf_richpresence_matchgroup_casual"		"Könnyed"
					"#tf_richpresence_matchgroup_specialevent"		"Különleges esemény"
					"#tf_richpresence_matchgroup_mannup"		"„Mannj rá”"
					"#tf_richpresence_matchgroup_bootcamp"		"Kiképzés"
				}
			}
			"italian"
			{
				"tokens"
				{
					"#tf_richpresence_display"		"{#TF_RichPresence_State_%state%}"
					"#tf_richpresence_state_mainmenu"		"Menu principale"
					"#tf_richpresence_state_searchinggeneric"		"In cerca di una partita"
					"#tf_richpresence_state_searchingmatchgroup"		"Ricerca in corso - {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playinggeneric"		"In partita - %currentmap%"
					"#tf_richpresence_state_loadinggeneric"		"Unione a una partita"
					"#tf_richpresence_state_playingmatchgroup"		"{#TF_RichPresence_MatchGroup_%matchgrouploc%} - %currentmap%"
					"#tf_richpresence_state_loadingmatchgroup"		"Unione a {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playingcommunity"		"Comunità - %currentmap%"
					"#tf_richpresence_state_loadingcommunity"		"Unione a un server della Comunità"
					"#tf_richpresence_matchgroup_competitive6v6"		"Competitiva"
					"#tf_richpresence_matchgroup_casual"		"Leggera"
					"#tf_richpresence_matchgroup_specialevent"		"Evento speciale"
					"#tf_richpresence_matchgroup_mannup"		"Mann Up"
					"#tf_richpresence_matchgroup_bootcamp"		"Addestramento"
				}
			}
			"polish"
			{
				"tokens"
				{
					"#tf_richpresence_display"		"{#TF_RichPresence_State_%state%}"
					"#tf_richpresence_state_mainmenu"		"Menu główne"
					"#tf_richpresence_state_searchinggeneric"		"Wyszukiwanie gry"
					"#tf_richpresence_state_searchingmatchgroup"		"Wyszukiwanie - {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playinggeneric"		"W grze - %currentmap%"
					"#tf_richpresence_state_loadinggeneric"		"Dołączanie do gry"
					"#tf_richpresence_state_playingmatchgroup"		"{#TF_RichPresence_MatchGroup_%matchgrouploc%} - %currentmap%"
					"#tf_richpresence_state_loadingmatchgroup"		"Dołączanie do trybu „{#TF_RichPresence_MatchGroup_%matchgrouploc%}”"
					"#tf_richpresence_state_playingcommunity"		"Gra społeczności - %currentmap%"
					"#tf_richpresence_state_loadingcommunity"		"Dołączanie do serwera społeczności"
					"#tf_richpresence_matchgroup_competitive6v6"		"Rankingowy"
					"#tf_richpresence_matchgroup_casual"		"Swobodny"
					"#tf_richpresence_matchgroup_specialevent"		"Wydarzenie specjalne"
					"#tf_richpresence_matchgroup_mannup"		"Mann Up"
					"#tf_richpresence_matchgroup_bootcamp"		"Trening"
				}
			}
			"russian"
			{
				"tokens"
				{
					"#tf_richpresence_display"		"{#TF_RichPresence_State_%state%}"
					"#tf_richpresence_state_mainmenu"		"В главном меню"
					"#tf_richpresence_state_searchinggeneric"		"Ищет игру"
					"#tf_richpresence_state_searchingmatchgroup"		"Поиск ({#TF_RichPresence_MatchGroup_%matchgrouploc%})"
					"#tf_richpresence_state_playinggeneric"		"В игре: %currentmap%"
					"#tf_richpresence_state_loadinggeneric"		"Заходит на сервер"
					"#tf_richpresence_state_playingmatchgroup"		"{#TF_RichPresence_MatchGroup_%matchgrouploc%}: %currentmap%"
					"#tf_richpresence_state_loadingmatchgroup"		"Заходит на сервер ({#TF_RichPresence_MatchGroup_%matchgrouploc%})"
					"#tf_richpresence_state_playingcommunity"		"Сервер сообщества: %currentmap%"
					"#tf_richpresence_state_loadingcommunity"		"Заходит на сервер сообщества"
					"#tf_richpresence_matchgroup_competitive6v6"		"Соревновательная игра"
					"#tf_richpresence_matchgroup_casual"		"Обычная игра"
					"#tf_richpresence_matchgroup_specialevent"		"Праздничный режим"
					"#tf_richpresence_matchgroup_mannup"		"МАННёвры"
					"#tf_richpresence_matchgroup_bootcamp"		"Учебка"
				}
			}
			"spanish"
			{
				"tokens"
				{
					"#tf_richpresence_display"		"{#TF_RichPresence_State_%state%}"
					"#tf_richpresence_state_mainmenu"		"Menú principal"
					"#tf_richpresence_state_searchinggeneric"		"Buscando una partida"
					"#tf_richpresence_state_searchingmatchgroup"		"Buscando - {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playinggeneric"		"En una partida - %currentmap%"
					"#tf_richpresence_state_loadinggeneric"		"Uniéndose a una partida"
					"#tf_richpresence_state_playingmatchgroup"		"{#TF_RichPresence_MatchGroup_%matchgrouploc%} - %currentmap%"
					"#tf_richpresence_state_loadingmatchgroup"		"Uniéndose a {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playingcommunity"		"Comunidad - %currentmap%"
					"#tf_richpresence_state_loadingcommunity"		"Uniéndose a un servidor de la comunidad"
					"#tf_richpresence_matchgroup_competitive6v6"		"Competitivo"
					"#tf_richpresence_matchgroup_casual"		"Casual"
					"#tf_richpresence_matchgroup_specialevent"		"Evento especial"
					"#tf_richpresence_matchgroup_mannup"		"Modo Mann"
					"#tf_richpresence_matchgroup_bootcamp"		"Iniciación"
				}
			}
			"thai"
			{
				"tokens"
				{
					"#tf_richpresence_display"		"{#TF_RichPresence_State_%state%}"
					"#tf_richpresence_state_mainmenu"		"เมนูหลัก"
					"#tf_richpresence_state_searchinggeneric"		"กำลังค้นหาแมตช์"
					"#tf_richpresence_state_searchingmatchgroup"		"กำลังค้นหา - {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playinggeneric"		"อยู่ในแมตช์ - %currentmap%"
					"#tf_richpresence_state_loadinggeneric"		"กำลังเข้าร่วมแมตช์"
					"#tf_richpresence_state_playingmatchgroup"		"{#TF_RichPresence_MatchGroup_%matchgrouploc%} - %currentmap%"
					"#tf_richpresence_state_loadingmatchgroup"		"กำลังเข้าร่วม {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playingcommunity"		"ชุมชน - %currentmap%"
					"#tf_richpresence_state_loadingcommunity"		"กำลังเข้าร่วมเซิร์ฟเวอร์ชุมชน"
					"#tf_richpresence_matchgroup_competitive6v6"		"การแข่งขัน"
					"#tf_richpresence_matchgroup_casual"		"แคชชวล"
					"#tf_richpresence_matchgroup_specialevent"		"กิจกรรมพิเศษ"
					"#tf_richpresence_matchgroup_mannup"		"Mann Up"
					"#tf_richpresence_matchgroup_bootcamp"		"บูทแคมป์"
				}
			}
			"turkish"
			{
				"tokens"
				{
					"#tf_richpresence_display"		"{#TF_RichPresence_State_%state%}"
					"#tf_richpresence_state_mainmenu"		"Ana Menü"
					"#tf_richpresence_state_searchinggeneric"		"Maç Aranıyor"
					"#tf_richpresence_state_searchingmatchgroup"		"Aranıyor - {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playinggeneric"		"Maçta - %currentmap%"
					"#tf_richpresence_state_loadinggeneric"		"Maça Katılınıyor"
					"#tf_richpresence_state_playingmatchgroup"		"{#TF_RichPresence_MatchGroup_%matchgrouploc%} - %currentmap%"
					"#tf_richpresence_state_loadingmatchgroup"		"Katılınıyor {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playingcommunity"		"Topluluk - %currentmap%"
					"#tf_richpresence_state_loadingcommunity"		"Topluluk Sunucusuna Katılınıyor"
					"#tf_richpresence_matchgroup_competitive6v6"		"Rekabetçi"
					"#tf_richpresence_matchgroup_casual"		"Basit Eğlence"
					"#tf_richpresence_matchgroup_specialevent"		"Özel Etkinlik"
					"#tf_richpresence_matchgroup_mannup"		"Mann Up"
					"#tf_richpresence_matchgroup_bootcamp"		"Eğitim"
				}
			}
			"brazilian"
			{
				"tokens"
				{
					"#tf_richpresence_display"		"{#TF_RichPresence_State_%state%}"
					"#tf_richpresence_state_mainmenu"		"Menu principal"
					"#tf_richpresence_state_searchinggeneric"		"Buscando partida"
					"#tf_richpresence_state_searchingmatchgroup"		"Buscando — {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playinggeneric"		"Em partida — %currentmap%"
					"#tf_richpresence_state_loadinggeneric"		"Entrando em partida"
					"#tf_richpresence_state_playingmatchgroup"		"{#TF_RichPresence_MatchGroup_%matchgrouploc%} — %currentmap%"
					"#tf_richpresence_state_loadingmatchgroup"		"Entrando em: {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playingcommunity"		"Servidor da comunidade — %currentmap%"
					"#tf_richpresence_state_loadingcommunity"		"Entrando em servidor da comunidade"
					"#tf_richpresence_matchgroup_competitive6v6"		"Competitivo"
					"#tf_richpresence_matchgroup_casual"		"Casual"
					"#tf_richpresence_matchgroup_specialevent"		"Evento especial"
					"#tf_richpresence_matchgroup_mannup"		"Mann Up"
					"#tf_richpresence_matchgroup_bootcamp"		"Treino (MvM)"
				}
			}
			"czech"
			{
				"tokens"
				{
					"#tf_richpresence_display"		"{#TF_RichPresence_State_%state%}"
					"#tf_richpresence_state_mainmenu"		"Hlavní menu"
					"#tf_richpresence_state_searchinggeneric"		"Vyhledávání zápasu"
					"#tf_richpresence_state_searchingmatchgroup"		"Vyhledávání – {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playinggeneric"		"Zápas – %currentmap%"
					"#tf_richpresence_state_loadinggeneric"		"Připojování do zápasu"
					"#tf_richpresence_state_playingmatchgroup"		"{#TF_RichPresence_MatchGroup_%matchgrouploc%} – %currentmap%"
					"#tf_richpresence_state_loadingmatchgroup"		"Připojování – {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playingcommunity"		"Komunitní server – %currentmap%"
					"#tf_richpresence_state_loadingcommunity"		"Připojování ke komunitnímu serveru"
					"#tf_richpresence_matchgroup_competitive6v6"		"Kompetitivní"
					"#tf_richpresence_matchgroup_casual"		"Nenáročný"
					"#tf_richpresence_matchgroup_specialevent"		"Speciální událost"
					"#tf_richpresence_matchgroup_mannup"		"Mann Up"
					"#tf_richpresence_matchgroup_bootcamp"		"Boot Camp"
				}
			}
			"finnish"
			{
				"tokens"
				{
					"#tf_richpresence_display"		"{#TF_RichPresence_State_%state%}"
					"#tf_richpresence_state_mainmenu"		"Päävalikko"
					"#tf_richpresence_state_searchinggeneric"		"Etsii peliä"
					"#tf_richpresence_state_searchingmatchgroup"		"Etsii - {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playinggeneric"		"Pelissä - %currentmap%"
					"#tf_richpresence_state_loadinggeneric"		"Liittymässä peliin"
					"#tf_richpresence_state_playingmatchgroup"		"{#TF_RichPresence_MatchGroup_%matchgrouploc%} - %currentmap%"
					"#tf_richpresence_state_loadingmatchgroup"		"Liittyy {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playingcommunity"		"Yhteisö - %currentmap%"
					"#tf_richpresence_state_loadingcommunity"		"Liittymässä yhteisöpalvelimelle"
					"#tf_richpresence_matchgroup_competitive6v6"		"Kilpailullinen"
					"#tf_richpresence_matchgroup_casual"		"Tavallinen"
					"#tf_richpresence_matchgroup_specialevent"		"Erikoistapahtuma"
					"#tf_richpresence_matchgroup_mannup"		"Miehisty-tila"
					"#tf_richpresence_matchgroup_bootcamp"		"Harjoitusleiri"
				}
			}
			"portuguese"
			{
				"tokens"
				{
					"#tf_richpresence_display"		"{#TF_RichPresence_State_%state%}"
					"#tf_richpresence_state_mainmenu"		"Menu principal"
					"#tf_richpresence_state_searchinggeneric"		"A procurar partida"
					"#tf_richpresence_state_searchingmatchgroup"		"A procurar - {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playinggeneric"		"Numa partida - %currentmap%"
					"#tf_richpresence_state_loadinggeneric"		"A entrar numa partida"
					"#tf_richpresence_state_playingmatchgroup"		"{#TF_RichPresence_MatchGroup_%matchgrouploc%} - %currentmap%"
					"#tf_richpresence_state_loadingmatchgroup"		"A entrar em: {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playingcommunity"		"Servidor comunitário - %currentmap%"
					"#tf_richpresence_state_loadingcommunity"		"A entrar num servidor comunitário"
					"#tf_richpresence_matchgroup_competitive6v6"		"Competitivo"
					"#tf_richpresence_matchgroup_casual"		"Casual"
					"#tf_richpresence_matchgroup_specialevent"		"Evento especial"
					"#tf_richpresence_matchgroup_mannup"		"Mann Up"
					"#tf_richpresence_matchgroup_bootcamp"		"Iniciação"
				}
			}
			"ukrainian"
			{
				"tokens"
				{
					"#tf_richpresence_display"		"{#TF_RichPresence_State_%state%}"
					"#tf_richpresence_state_mainmenu"		"Головне меню"
					"#tf_richpresence_state_searchinggeneric"		"Пошук матчу"
					"#tf_richpresence_state_searchingmatchgroup"		"Пошук — {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playinggeneric"		"У матчі — %currentmap%"
					"#tf_richpresence_state_loadinggeneric"		"Приєднується до матчу"
					"#tf_richpresence_state_playingmatchgroup"		"{#TF_RichPresence_MatchGroup_%matchgrouploc%} — %currentmap%"
					"#tf_richpresence_state_loadingmatchgroup"		"Приєднується до {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playingcommunity"		"Спільнота — %currentmap%"
					"#tf_richpresence_state_loadingcommunity"		"Приєднується до сервера спільноти"
					"#tf_richpresence_matchgroup_competitive6v6"		"Змагальний режим"
					"#tf_richpresence_matchgroup_casual"		"Звичайний режим"
					"#tf_richpresence_matchgroup_specialevent"		"Особлива подія"
					"#tf_richpresence_matchgroup_mannup"		"«Манневри»"
					"#tf_richpresence_matchgroup_bootcamp"		"«Тренувальний табір»"
				}
			}
			"swedish"
			{
				"tokens"
				{
					"#tf_richpresence_display"		"{#TF_RichPresence_State_%state%}"
					"#tf_richpresence_state_mainmenu"		"Huvudmeny"
					"#tf_richpresence_state_searchinggeneric"		"Söker efter en match"
					"#tf_richpresence_state_searchingmatchgroup"		"Söker - {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playinggeneric"		"I match - %currentmap%"
					"#tf_richpresence_state_loadinggeneric"		"Går med i match"
					"#tf_richpresence_state_playingmatchgroup"		"{#TF_RichPresence_MatchGroup_%matchgrouploc%} - %currentmap%"
					"#tf_richpresence_state_loadingmatchgroup"		"Går med i {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playingcommunity"		"Gemenskap - %currentmap%"
					"#tf_richpresence_state_loadingcommunity"		"Går med i gemenskapserver"
					"#tf_richpresence_matchgroup_competitive6v6"		"Tävlingsinriktat"
					"#tf_richpresence_matchgroup_casual"		"Vanligt"
					"#tf_richpresence_matchgroup_specialevent"		"Specialevent"
					"#tf_richpresence_matchgroup_mannup"		"Mann Up"
					"#tf_richpresence_matchgroup_bootcamp"		"Träningsläger"
				}
			}
			"greek"
			{
				"tokens"
				{
					"#tf_richpresence_display"		"{#TF_RichPresence_State_%state%}"
					"#tf_richpresence_state_mainmenu"		"Κύριο μενού"
					"#tf_richpresence_state_searchinggeneric"		"Αναζήτηση αγώνα"
					"#tf_richpresence_state_searchingmatchgroup"		"Αναζήτηση - {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playinggeneric"		"Σε αγώνα - %currentmap%"
					"#tf_richpresence_state_loadinggeneric"		"Σύνδεση σε αγώνα"
					"#tf_richpresence_state_playingmatchgroup"		"{#TF_RichPresence_MatchGroup_%matchgrouploc%} - %currentmap%"
					"#tf_richpresence_state_loadingmatchgroup"		"Σύνδεση σε {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playingcommunity"		"Κοινότητα - %currentmap%"
					"#tf_richpresence_state_loadingcommunity"		"Σύνδεση σε διακομιστή κοινότητας"
					"#tf_richpresence_matchgroup_competitive6v6"		"Competitive"
					"#tf_richpresence_matchgroup_casual"		"Casual"
					"#tf_richpresence_matchgroup_specialevent"		"Ειδικό συμβάν"
					"#tf_richpresence_matchgroup_mannup"		"Mann Up"
					"#tf_richpresence_matchgroup_bootcamp"		"Boot Camp"
				}
			}
			"norwegian"
			{
				"tokens"
				{
					"#tf_richpresence_display"		"{#TF_RichPresence_State_%state%}"
					"#tf_richpresence_state_mainmenu"		"Hovedmeny"
					"#tf_richpresence_state_searchinggeneric"		"Søker etter en kamp"
					"#tf_richpresence_state_searchingmatchgroup"		"Søker – {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playinggeneric"		"I kamp – %currentmap%"
					"#tf_richpresence_state_loadinggeneric"		"Blir med i en kamp"
					"#tf_richpresence_state_playingmatchgroup"		"{#TF_RichPresence_MatchGroup_%matchgrouploc%} – %currentmap%"
					"#tf_richpresence_state_loadingmatchgroup"		"Blir med i {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playingcommunity"		"Samfunn – %currentmap%"
					"#tf_richpresence_state_loadingcommunity"		"Blir med på samfunnstjener"
					"#tf_richpresence_matchgroup_competitive6v6"		"Konkurransespilling"
					"#tf_richpresence_matchgroup_casual"		"Avslappet"
					"#tf_richpresence_matchgroup_specialevent"		"Spesielt arrangement"
					"#tf_richpresence_matchgroup_mannup"		"Mann deg opp"
					"#tf_richpresence_matchgroup_bootcamp"		"Treningsleir"
				}
			}
			"schinese"
			{
				"tokens"
				{
					"#tf_richpresence_display"		"{#TF_RichPresence_State_%state%}"
					"#tf_richpresence_state_mainmenu"		"主菜单"
					"#tf_richpresence_state_searchinggeneric"		"正在搜索比赛"
					"#tf_richpresence_state_searchingmatchgroup"		"正在搜索 - {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playinggeneric"		"比赛中 - %currentmap%"
					"#tf_richpresence_state_loadinggeneric"		"正在加入比赛"
					"#tf_richpresence_state_playingmatchgroup"		"{#TF_RichPresence_MatchGroup_%matchgrouploc%} - %currentmap%"
					"#tf_richpresence_state_loadingmatchgroup"		"正在加入 {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playingcommunity"		"社区服务器 - %currentmap%"
					"#tf_richpresence_state_loadingcommunity"		"正在加入社区服务器"
					"#tf_richpresence_matchgroup_competitive6v6"		"竞技模式"
					"#tf_richpresence_matchgroup_casual"		"休闲模式"
					"#tf_richpresence_matchgroup_specialevent"		"特别活动"
					"#tf_richpresence_matchgroup_mannup"		"曼恩奇现"
					"#tf_richpresence_matchgroup_bootcamp"		"新兵训练营"
				}
			}
			"dutch"
			{
				"tokens"
				{
					"#tf_richpresence_display"		"{#TF_RichPresence_State_%state%}"
					"#tf_richpresence_state_mainmenu"		"Hoofdmenu"
					"#tf_richpresence_state_searchinggeneric"		"Naar een spel aan het zoeken"
					"#tf_richpresence_state_searchingmatchgroup"		"Zoeken - {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playinggeneric"		"In spel - %currentmap%"
					"#tf_richpresence_state_playingmatchgroup"		"{#TF_RichPresence_MatchGroup_%matchgrouploc%} - %currentmap%"
					"#tf_richpresence_state_playingcommunity"		"Community - %currentmap%"
					"#tf_richpresence_matchgroup_competitive6v6"		"Competitief"
					"#tf_richpresence_matchgroup_casual"		"Casual"
					"#tf_richpresence_matchgroup_specialevent"		"Speciaal evenement"
					"#tf_richpresence_matchgroup_mannup"		"Vermann je"
					"#tf_richpresence_matchgroup_bootcamp"		"Trainingskamp"
				}
			}
			"japanese"
			{
				"tokens"
				{
					"#tf_richpresence_state_mainmenu"		"メインメニュー"
					"#tf_richpresence_state_loadinggeneric"		"マッチに参加しています"
					"#tf_richpresence_matchgroup_competitive6v6"		"対戦"
					"#tf_richpresence_matchgroup_casual"		"カジュアル"
					"#tf_richpresence_matchgroup_specialevent"		"スペシャルイベント"
					"#tf_richpresence_matchgroup_bootcamp"		"ブート・キャンプ"
				}
			}
			"tchinese"
			{
				"tokens"
				{
					"#tf_richpresence_display"		"{#TF_RichPresence_State_%state%}"
					"#tf_richpresence_state_mainmenu"		"主選單"
					"#tf_richpresence_state_searchinggeneric"		"搜尋比賽中"
					"#tf_richpresence_state_searchingmatchgroup"		"搜尋中 - {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playinggeneric"		"比賽中 - %currentmap%"
					"#tf_richpresence_state_loadinggeneric"		"加入比賽中"
					"#tf_richpresence_state_playingmatchgroup"		"{#TF_RichPresence_MatchGroup_%matchgrouploc%} - %currentmap%"
					"#tf_richpresence_state_loadingmatchgroup"		"加入 {#TF_RichPresence_MatchGroup_%matchgrouploc%} 中"
					"#tf_richpresence_state_playingcommunity"		"社群 - %currentmap%"
					"#tf_richpresence_state_loadingcommunity"		"加入社群伺服器中"
					"#tf_richpresence_matchgroup_competitive6v6"		"競技模式"
					"#tf_richpresence_matchgroup_casual"		"休閒模式"
					"#tf_richpresence_matchgroup_specialevent"		"特殊活動"
					"#tf_richpresence_matchgroup_mannup"		"曼恩對決機器曼起來模式"
					"#tf_richpresence_matchgroup_bootcamp"		"曼恩對決機器訓練營"
				}
			}
			"vietnamese"
			{
				"tokens"
				{
					"#tf_richpresence_display"		"{#TF_RichPresence_State_%state%}"
					"#tf_richpresence_state_mainmenu"		"Đại sảnh"
					"#tf_richpresence_state_searchinggeneric"		"Đang tìm trận đấu"
					"#tf_richpresence_state_searchingmatchgroup"		"Đang tìm - {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playinggeneric"		"Đang chơi - %currentmap%"
					"#tf_richpresence_state_loadinggeneric"		"Đang tham gia trận"
					"#tf_richpresence_state_playingmatchgroup"		"{#TF_RichPresence_MatchGroup_%matchgrouploc%} - %currentmap%"
					"#tf_richpresence_state_loadingmatchgroup"		"Đang tham gia {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playingcommunity"		"Cộng đồng - %currentmap%"
					"#tf_richpresence_state_loadingcommunity"		"Đang tham gia máy chủ cộng đồng"
					"#tf_richpresence_matchgroup_competitive6v6"		"Tranh đấu"
					"#tf_richpresence_matchgroup_casual"		"Đơn giản"
					"#tf_richpresence_matchgroup_specialevent"		"Sự kiện đặc biệt"
					"#tf_richpresence_matchgroup_mannup"		"Mann Up"
					"#tf_richpresence_matchgroup_bootcamp"		"Boot Camp"
				}
			}
			"latam"
			{
				"tokens"
				{
					"#tf_richpresence_display"		"{#TF_RichPresence_State_%state%}"
					"#tf_richpresence_state_mainmenu"		"Menú principal"
					"#tf_richpresence_state_searchinggeneric"		"Buscando una partida"
					"#tf_richpresence_state_searchingmatchgroup"		"Buscando - {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playinggeneric"		"En una partida - %currentmap%"
					"#tf_richpresence_state_loadinggeneric"		"Uniéndose a una partida"
					"#tf_richpresence_state_playingmatchgroup"		"{#TF_RichPresence_MatchGroup_%matchgrouploc%} - %currentmap%"
					"#tf_richpresence_state_loadingmatchgroup"		"Uniéndose a {#TF_RichPresence_MatchGroup_%matchgrouploc%}"
					"#tf_richpresence_state_playingcommunity"		"Comunidad - %currentmap%"
					"#tf_richpresence_state_loadingcommunity"		"Uniéndose a un servidor de la comunidad"
					"#tf_richpresence_matchgroup_competitive6v6"		"Competitivo"
					"#tf_richpresence_matchgroup_casual"		"Casual"
					"#tf_richpresence_matchgroup_specialevent"		"Evento especial"
					"#tf_richpresence_matchgroup_mannup"		"Modo Mann"
					"#tf_richpresence_matchgroup_bootcamp"		"Iniciación"
				}
			}
		}
	}
}`)

func TestParseResponse(t *testing.T) {
	steamJson := SteamResponseToJSON(testSteamPayload)
	var manifest RemoteManifest
	err := json.Unmarshal(steamJson, &manifest)
	assert.NoError(t, err)
	assert.Equal(t, "14136189", manifest.GetBuildIDForBranch("public"))
}
