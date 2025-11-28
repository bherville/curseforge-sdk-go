# CurseForge SDK for Go

A Go client library for the [CurseForge API](https://docs.curseforge.com/). This library provides easy access to CurseForge's API for searching mods, retrieving mod information, checking for updates, and more.

## Installation

```bash
go get github.com/bherville/curseforge-sdk-go
```

## Requirements

- Go 1.20 or later
- CurseForge API key (obtain from [CurseForge Console](https://console.curseforge.com/))

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/bherville/curseforge-sdk-go/pkg/curseforge"
)

func main() {
    // Create a server configuration with your API key
    server := curseforge.NewServer("your-api-key")

    // Search for mods
    mods, _, err := curseforge.SearchMods(server, curseforge.SearchModsRequest{
        GameID:       curseforge.GameIDMinecraft,
        SearchFilter: "jei",
        ClassID:      curseforge.ClassIDMods,
    })
    if err != nil {
        panic(err)
    }

    for _, mod := range mods {
        fmt.Printf("Found: %s (ID: %d)\n", mod.Name, mod.ID)
    }
}
```

## Features

### Search Mods

```go
mods, pagination, err := curseforge.SearchMods(server, curseforge.SearchModsRequest{
    GameID:        curseforge.GameIDMinecraft,
    ClassID:       curseforge.ClassIDMods,
    SearchFilter:  "optimization",
    GameVersion:   "1.20.1",
    ModLoaderType: curseforge.ModLoaderFabric,
    SortField:     curseforge.SortFieldPopularity,
    SortOrder:     "desc",
    PageSize:      20,
})
```

### Get Mod Details

```go
// Get a specific mod by ID
mod, err := curseforge.GetMod(server, 238222) // JEI

// Get multiple mods
mods, err := curseforge.GetMods(server, []int{238222, 306612})

// Get mod description (HTML)
description, err := curseforge.GetModDescription(server, 238222)
```

### Get Mod Files

```go
// Get all files for a mod
files, pagination, err := curseforge.GetModFiles(server, modID, nil)

// Get files with filtering
files, pagination, err := curseforge.GetModFiles(server, modID, &curseforge.GetModFilesRequest{
    GameVersion:   "1.20.1",
    ModLoaderType: curseforge.ModLoaderFabric,
    PageSize:      50,
})

// Get a specific file
file, err := curseforge.GetModFile(server, modID, fileID)

// Get file changelog
changelog, err := curseforge.GetModFileChangelog(server, modID, fileID)

// Get download URL
url, err := curseforge.GetModFileDownloadURL(server, modID, fileID)
```

### Fingerprint Matching

CurseForge uses Murmur2 hashes (fingerprints) to identify mod files:

```go
// Compute fingerprint for a local file
fingerprint, err := curseforge.ComputeFileFingerprint("/path/to/mod.jar")

// Match fingerprints against CurseForge database
matches, err := curseforge.GetFingerprintsMatches(server, []int64{fingerprint})

if len(matches.ExactMatches) > 0 {
    match := matches.ExactMatches[0]
    fmt.Printf("Found mod: %s\n", match.File.DisplayName)
}

// Check unmatched fingerprints
for _, fp := range matches.UnmatchedFingerprints {
    fmt.Printf("No match for fingerprint: %d\n", fp)
}
```

### Download Files

```go
file, _ := curseforge.GetModFile(server, modID, fileID)
err := curseforge.DownloadFile(file, "/path/to/destination.jar")
```

### Minecraft-Specific APIs

```go
// Get all Minecraft versions
versions, err := curseforge.GetMinecraftVersions(server)

// Get specific version info
versionInfo, err := curseforge.GetSpecificMinecraftVersion(server, "1.20.1")

// Get all mod loaders
loaders, err := curseforge.GetMinecraftModLoaders(server)

// Get mod loaders for a specific Minecraft version
loaders, err := curseforge.GetMinecraftModLoadersForVersion(server, "1.20.1")
```

### Categories

```go
// Get all categories for Minecraft
categories, err := curseforge.GetCategories(server, curseforge.GameIDMinecraft)

// Get categories for a specific class (e.g., mods)
categories, err := curseforge.GetCategoriesByClassID(server, curseforge.GameIDMinecraft, curseforge.ClassIDMods)
```

## Constants

### Game IDs

```go
curseforge.GameIDMinecraft // 432
```

### Class IDs

```go
curseforge.ClassIDMods          // 6
curseforge.ClassIDModpacks      // 4471
curseforge.ClassIDResourcePacks // 12
curseforge.ClassIDWorlds        // 17
curseforge.ClassIDShaders       // 6552
```

### Mod Loaders

```go
curseforge.ModLoaderAny       // 0
curseforge.ModLoaderForge     // 1
curseforge.ModLoaderFabric    // 4
curseforge.ModLoaderQuilt     // 5
curseforge.ModLoaderNeoForge  // 6
```

### Sort Fields

```go
curseforge.SortFieldFeatured       // 1
curseforge.SortFieldPopularity     // 2
curseforge.SortFieldLastUpdated    // 3
curseforge.SortFieldName           // 4
curseforge.SortFieldTotalDownloads // 6
curseforge.SortFieldRating         // 12
```

### Release Types

```go
curseforge.ReleaseTypeRelease // 1
curseforge.ReleaseTypeBeta    // 2
curseforge.ReleaseTypeAlpha   // 3
```

## Helper Functions

### Hash Extraction

```go
sha1 := curseforge.GetSha1Hash(file)
md5 := curseforge.GetMd5Hash(file)
```

### Version Checking

```go
if file.HasGameVersion("1.20.1") {
    fmt.Println("Supports 1.20.1")
}

if file.HasModLoader(curseforge.ModLoaderFabric) {
    fmt.Println("Supports Fabric")
}
```

## Debug Logging

Enable trace logging to see API requests and responses:

```go
import log "github.com/sirupsen/logrus"

log.SetLevel(log.TraceLevel)
```

## API Reference

For complete API documentation, see the [CurseForge API Docs](https://docs.curseforge.com/).

## License

MIT License - see [LICENSE](LICENSE) file.
