// Package curseforge provides a Go client library for the CurseForge API.
//
// This library allows you to interact with the CurseForge API to search for mods,
// retrieve mod information, check for updates, and more. It is designed primarily
// for Minecraft mod management but supports other games available on CurseForge.
//
// # Authentication
//
// CurseForge API requires an API key for all requests. You can obtain an API key
// from the CurseForge console: https://console.curseforge.com/
//
// # Quick Start
//
// Create a server configuration with your API key:
//
//server := curseforge.NewServer("your-api-key")
//
// Search for mods:
//
//mods, pagination, err := curseforge.SearchMods(server, curseforge.SearchModsRequest{
//    GameID:       curseforge.GameIDMinecraft,
//    SearchFilter: "jei",
//    ClassID:      curseforge.ClassIDMods,
//})
//
// Get a specific mod:
//
//mod, err := curseforge.GetMod(server, 238222) // JEI mod ID
//
// Get mod files with filtering:
//
//files, pagination, err := curseforge.GetModFiles(server, modID, &curseforge.GetModFilesRequest{
//    GameVersion:   "1.20.1",
//    ModLoaderType: curseforge.ModLoaderFabric,
//})
//
// # Fingerprint Matching
//
// CurseForge uses Murmur2 hashes (fingerprints) to identify mod files. You can
// compute fingerprints for local files and match them against the CurseForge database:
//
//fingerprint, err := curseforge.ComputeFileFingerprint("/path/to/mod.jar")
//matches, err := curseforge.GetFingerprintsMatches(server, []int64{fingerprint})
//
// # Mod Loaders
//
// The library supports various mod loaders:
//
//   - ModLoaderForge - Minecraft Forge
//   - ModLoaderFabric - Fabric
//   - ModLoaderQuilt - Quilt
//   - ModLoaderNeoForge - NeoForge
//   - ModLoaderLiteLoader - LiteLoader (legacy)
//
// # Game and Class IDs
//
// Common constants are provided:
//
//   - GameIDMinecraft (432) - Minecraft game ID
//   - ClassIDMods (6) - Mods category
//   - ClassIDModpacks (4471) - Modpacks category
//   - ClassIDResourcePacks (12) - Resource packs category
//
// # Error Handling
//
// All API functions return errors that should be checked. API errors include
// the error message and description from CurseForge.
//
// # Logging
//
// The library uses logrus for logging. Set the log level to trace to see
// detailed API request/response information:
//
//import log "github.com/sirupsen/logrus"
//log.SetLevel(log.TraceLevel)
package curseforge
