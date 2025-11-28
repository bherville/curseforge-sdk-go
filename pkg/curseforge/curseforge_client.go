package curseforge

import (
"bytes"
"encoding/json"
"fmt"
"io"
"net/http"
"net/url"
"os"
"strconv"

log "github.com/sirupsen/logrus"
)

var (
contextLogger *log.Entry
)

const (
// API Endpoints
ApiEndpointGames             = "v1/games"
ApiEndpointMods              = "v1/mods"
ApiEndpointFingerprints      = "v1/fingerprints"
ApiEndpointFingerprintsFuzzy = "v1/fingerprints/fuzzy"
ApiEndpointCategories        = "v1/categories"
ApiEndpointMinecraft         = "v1/minecraft"

// Base URLs
BaseURLProduction = "https://api.curseforge.com"
)

func init() {
contextLogger = log.WithFields(log.Fields{
"library": "curseforge-sdk-go",
})
}

// NewServer creates a new CurseForge server configuration
// apiKey is required for CurseForge API access
func NewServer(apiKey string) CurseForgeServer {
return CurseForgeServer{
Name:   "Production",
Url:    BaseURLProduction,
ApiKey: apiKey,
}
}

// NewServerWithURL creates a new CurseForge server with a custom URL
func NewServerWithURL(apiKey string, baseURL string) CurseForgeServer {
return CurseForgeServer{
Name:   "Custom",
Url:    baseURL,
ApiKey: apiKey,
}
}

func buildApiUrl(server CurseForgeServer, endpoint string, subPaths []string) string {
apiUrl := fmt.Sprintf("%s/%s", server.Url, endpoint)

for _, path := range subPaths {
apiUrl = fmt.Sprintf("%s/%s", apiUrl, path)
}
return apiUrl
}

func callApi[T any](apiObject *T, server CurseForgeServer, method string, endpoint string, subPaths []string, queryParams map[string]string, body interface{}) error {
apiUrl := buildApiUrl(server, endpoint, subPaths)

// Add query parameters
if queryParams != nil && len(queryParams) > 0 {
params := url.Values{}
for k, v := range queryParams {
params.Set(k, v)
}
apiUrl = fmt.Sprintf("%s?%s", apiUrl, params.Encode())
}

contextLogger.Trace(fmt.Sprintf("apiUrl: %s", apiUrl))

var reqBody io.Reader
if body != nil {
jsonBody, err := json.Marshal(body)
if err != nil {
return fmt.Errorf("failed to marshal request body: %w", err)
}
reqBody = bytes.NewBuffer(jsonBody)
contextLogger.Trace(fmt.Sprintf("Request Body: %s", string(jsonBody)))
}

req, err := http.NewRequest(method, apiUrl, reqBody)
if err != nil {
return fmt.Errorf("failed to create request: %w", err)
}

req.Header.Add("Accept", "application/json")
req.Header.Add("x-api-key", server.ApiKey)
if body != nil {
req.Header.Add("Content-Type", "application/json")
}

res, err := http.DefaultClient.Do(req)
if err != nil {
return err
}

defer res.Body.Close()
resBody, err := io.ReadAll(res.Body)
if err != nil {
return fmt.Errorf("failed to read response body: %w", err)
}

contextLogger.Trace(fmt.Sprintf("Response Status: %d", res.StatusCode))
contextLogger.Trace(fmt.Sprintf("Response Body: %s", string(resBody)))

if res.StatusCode != http.StatusOK {
var apiError ApiError
if err := json.Unmarshal(resBody, &apiError); err != nil {
return fmt.Errorf("API request failed with status %d: %s", res.StatusCode, string(resBody))
}
return fmt.Errorf("API call failed: %s - %s", apiError.Error, apiError.Description)
}

if err := json.Unmarshal(resBody, &apiObject); err != nil {
return fmt.Errorf("failed to unmarshal response: %w", err)
}

return nil
}

// ============================================================================
// Games API
// ============================================================================

// GetGames retrieves all games available on CurseForge
func GetGames(server CurseForgeServer) ([]Game, error) {
var response PaginatedResponse[[]Game]
err := callApi(&response, server, http.MethodGet, ApiEndpointGames, nil, nil, nil)
if err != nil {
return nil, err
}
return response.Data, nil
}

// GetGame retrieves a specific game by ID
func GetGame(server CurseForgeServer, gameID int) (*Game, error) {
var response Response[Game]
err := callApi(&response, server, http.MethodGet, ApiEndpointGames, []string{strconv.Itoa(gameID)}, nil, nil)
if err != nil {
return nil, err
}
return &response.Data, nil
}

// GetGameVersions retrieves game versions for a specific game
func GetGameVersions(server CurseForgeServer, gameID int) ([]GameVersionType, error) {
var response Response[[]GameVersionType]
err := callApi(&response, server, http.MethodGet, ApiEndpointGames, []string{strconv.Itoa(gameID), "versions"}, nil, nil)
if err != nil {
return nil, err
}
return response.Data, nil
}

// GetGameVersionTypes retrieves game version types for a specific game
func GetGameVersionTypes(server CurseForgeServer, gameID int) ([]GameVersionType, error) {
var response Response[[]GameVersionType]
err := callApi(&response, server, http.MethodGet, ApiEndpointGames, []string{strconv.Itoa(gameID), "version-types"}, nil, nil)
if err != nil {
return nil, err
}
return response.Data, nil
}

// ============================================================================
// Mods API
// ============================================================================

// SearchMods searches for mods matching the given criteria
func SearchMods(server CurseForgeServer, request SearchModsRequest) ([]Mod, *Pagination, error) {
var response PaginatedResponse[[]Mod]

params := make(map[string]string)
params["gameId"] = strconv.Itoa(request.GameID)

if request.ClassID != 0 {
params["classId"] = strconv.Itoa(request.ClassID)
}
if request.CategoryID != 0 {
params["categoryId"] = strconv.Itoa(request.CategoryID)
}
if request.GameVersion != "" {
params["gameVersion"] = request.GameVersion
}
if request.SearchFilter != "" {
params["searchFilter"] = request.SearchFilter
}
if request.SortField != 0 {
params["sortField"] = strconv.Itoa(request.SortField)
}
if request.SortOrder != "" {
params["sortOrder"] = request.SortOrder
}
if request.ModLoaderType != ModLoaderAny {
params["modLoaderType"] = strconv.Itoa(int(request.ModLoaderType))
}
if request.GameVersionTypeID != 0 {
params["gameVersionTypeId"] = strconv.Itoa(request.GameVersionTypeID)
}
if request.AuthorID != 0 {
params["authorId"] = strconv.Itoa(request.AuthorID)
}
if request.Slug != "" {
params["slug"] = request.Slug
}
if request.Index != 0 {
params["index"] = strconv.Itoa(request.Index)
}
if request.PageSize != 0 {
params["pageSize"] = strconv.Itoa(request.PageSize)
}

err := callApi(&response, server, http.MethodGet, ApiEndpointMods, []string{"search"}, params, nil)
if err != nil {
return nil, nil, err
}

return response.Data, &response.Pagination, nil
}

// GetMod retrieves a specific mod by ID
func GetMod(server CurseForgeServer, modID int) (*Mod, error) {
var response Response[Mod]
err := callApi(&response, server, http.MethodGet, ApiEndpointMods, []string{strconv.Itoa(modID)}, nil, nil)
if err != nil {
return nil, err
}
return &response.Data, nil
}

// GetMods retrieves multiple mods by their IDs
func GetMods(server CurseForgeServer, modIDs []int) ([]Mod, error) {
var response Response[[]Mod]
body := GetModsByIDsRequest{ModIDs: modIDs}
err := callApi(&response, server, http.MethodPost, ApiEndpointMods, nil, nil, body)
if err != nil {
return nil, err
}
return response.Data, nil
}

// GetModDescription retrieves the description of a mod
func GetModDescription(server CurseForgeServer, modID int) (string, error) {
var response Response[string]
err := callApi(&response, server, http.MethodGet, ApiEndpointMods, []string{strconv.Itoa(modID), "description"}, nil, nil)
if err != nil {
return "", err
}
return response.Data, nil
}

// ============================================================================
// Files API
// ============================================================================

// GetModFile retrieves a specific file for a mod
func GetModFile(server CurseForgeServer, modID int, fileID int) (*File, error) {
var response Response[File]
err := callApi(&response, server, http.MethodGet, ApiEndpointMods, []string{strconv.Itoa(modID), "files", strconv.Itoa(fileID)}, nil, nil)
if err != nil {
return nil, err
}
return &response.Data, nil
}

// GetModFiles retrieves all files for a mod with optional filtering
func GetModFiles(server CurseForgeServer, modID int, request *GetModFilesRequest) ([]File, *Pagination, error) {
var response PaginatedResponse[[]File]

params := make(map[string]string)
if request != nil {
if request.GameVersion != "" {
params["gameVersion"] = request.GameVersion
}
if request.ModLoaderType != ModLoaderAny {
params["modLoaderType"] = strconv.Itoa(int(request.ModLoaderType))
}
if request.GameVersionTypeID != 0 {
params["gameVersionTypeId"] = strconv.Itoa(request.GameVersionTypeID)
}
if request.Index != 0 {
params["index"] = strconv.Itoa(request.Index)
}
if request.PageSize != 0 {
params["pageSize"] = strconv.Itoa(request.PageSize)
}
}

err := callApi(&response, server, http.MethodGet, ApiEndpointMods, []string{strconv.Itoa(modID), "files"}, params, nil)
if err != nil {
return nil, nil, err
}

return response.Data, &response.Pagination, nil
}

// GetFiles retrieves multiple files by their IDs
func GetFiles(server CurseForgeServer, fileIDs []int) ([]File, error) {
var response Response[[]File]
body := GetFilesRequest{FileIDs: fileIDs}
err := callApi(&response, server, http.MethodPost, ApiEndpointMods, []string{"files"}, nil, body)
if err != nil {
return nil, err
}
return response.Data, nil
}

// GetModFileChangelog retrieves the changelog for a specific file
func GetModFileChangelog(server CurseForgeServer, modID int, fileID int) (string, error) {
var response Response[string]
err := callApi(&response, server, http.MethodGet, ApiEndpointMods, []string{strconv.Itoa(modID), "files", strconv.Itoa(fileID), "changelog"}, nil, nil)
if err != nil {
return "", err
}
return response.Data, nil
}

// GetModFileDownloadURL retrieves the download URL for a specific file
func GetModFileDownloadURL(server CurseForgeServer, modID int, fileID int) (string, error) {
var response Response[string]
err := callApi(&response, server, http.MethodGet, ApiEndpointMods, []string{strconv.Itoa(modID), "files", strconv.Itoa(fileID), "download-url"}, nil, nil)
if err != nil {
return "", err
}
return response.Data, nil
}

// ============================================================================
// Fingerprints API
// ============================================================================

// GetFingerprintsMatches finds mods matching the given file fingerprints (Murmur2 hashes)
func GetFingerprintsMatches(server CurseForgeServer, fingerprints []int64) (*FingerprintMatchesResult, error) {
var response Response[FingerprintMatchesResult]
body := FingerprintsMatchesRequest{Fingerprints: fingerprints}
err := callApi(&response, server, http.MethodPost, ApiEndpointFingerprints, nil, nil, body)
if err != nil {
return nil, err
}
return &response.Data, nil
}

// GetFingerprintsMatchesByGameID finds mods matching fingerprints for a specific game
func GetFingerprintsMatchesByGameID(server CurseForgeServer, gameID int, fingerprints []int64) (*FingerprintMatchesResult, error) {
var response Response[FingerprintMatchesResult]
body := FingerprintsMatchesRequest{Fingerprints: fingerprints}
err := callApi(&response, server, http.MethodPost, ApiEndpointFingerprints, []string{strconv.Itoa(gameID)}, nil, body)
if err != nil {
return nil, err
}
return &response.Data, nil
}

// ============================================================================
// Categories API
// ============================================================================

// GetCategories retrieves all categories for a game
func GetCategories(server CurseForgeServer, gameID int) ([]Category, error) {
var response Response[[]Category]
params := map[string]string{"gameId": strconv.Itoa(gameID)}
err := callApi(&response, server, http.MethodGet, ApiEndpointCategories, nil, params, nil)
if err != nil {
return nil, err
}
return response.Data, nil
}

// GetCategoriesByClassID retrieves categories for a specific class
func GetCategoriesByClassID(server CurseForgeServer, gameID int, classID int) ([]Category, error) {
var response Response[[]Category]
params := map[string]string{
"gameId":  strconv.Itoa(gameID),
"classId": strconv.Itoa(classID),
}
err := callApi(&response, server, http.MethodGet, ApiEndpointCategories, nil, params, nil)
if err != nil {
return nil, err
}
return response.Data, nil
}

// ============================================================================
// Minecraft-specific API
// ============================================================================

// GetMinecraftVersions retrieves all Minecraft versions
func GetMinecraftVersions(server CurseForgeServer) ([]MinecraftVersionInfo, error) {
var response Response[[]MinecraftVersionInfo]
err := callApi(&response, server, http.MethodGet, ApiEndpointMinecraft, []string{"version"}, nil, nil)
if err != nil {
return nil, err
}
return response.Data, nil
}

// GetSpecificMinecraftVersion retrieves info for a specific Minecraft version
func GetSpecificMinecraftVersion(server CurseForgeServer, gameVersionString string) (*MinecraftVersionInfo, error) {
var response Response[MinecraftVersionInfo]
err := callApi(&response, server, http.MethodGet, ApiEndpointMinecraft, []string{"version", gameVersionString}, nil, nil)
if err != nil {
return nil, err
}
return &response.Data, nil
}

// GetMinecraftModLoaders retrieves all Minecraft mod loaders
func GetMinecraftModLoaders(server CurseForgeServer) ([]MinecraftModLoaderInfo, error) {
var response Response[[]MinecraftModLoaderInfo]
err := callApi(&response, server, http.MethodGet, ApiEndpointMinecraft, []string{"modloader"}, nil, nil)
if err != nil {
return nil, err
}
return response.Data, nil
}

// GetMinecraftModLoadersForVersion retrieves mod loaders for a specific Minecraft version
func GetMinecraftModLoadersForVersion(server CurseForgeServer, version string) ([]MinecraftModLoaderInfo, error) {
var response Response[[]MinecraftModLoaderInfo]
params := map[string]string{"version": version}
err := callApi(&response, server, http.MethodGet, ApiEndpointMinecraft, []string{"modloader"}, params, nil)
if err != nil {
return nil, err
}
return response.Data, nil
}

// GetSpecificMinecraftModLoader retrieves info for a specific mod loader
func GetSpecificMinecraftModLoader(server CurseForgeServer, modLoaderName string) (*MinecraftModLoaderInfo, error) {
var response Response[MinecraftModLoaderInfo]
err := callApi(&response, server, http.MethodGet, ApiEndpointMinecraft, []string{"modloader", modLoaderName}, nil, nil)
if err != nil {
return nil, err
}
return &response.Data, nil
}

// ============================================================================
// Download Helper
// ============================================================================

// DownloadFile downloads a mod file to the specified destination
func DownloadFile(file File, destination string) error {
if file.DownloadURL == "" {
return fmt.Errorf("download URL is not available for this file")
}

out, err := os.Create(destination)
if err != nil {
return err
}
defer out.Close()

resp, err := http.Get(file.DownloadURL)
if err != nil {
return err
}
defer resp.Body.Close()

if resp.StatusCode != http.StatusOK {
return fmt.Errorf("download failed with status: %d", resp.StatusCode)
}

_, err = io.Copy(out, resp.Body)
return err
}

// ============================================================================
// Fingerprint Helpers
// ============================================================================

// GetSha1Hash returns the SHA1 hash from a file's hashes if available
func GetSha1Hash(file File) string {
for _, hash := range file.Hashes {
if hash.Algo == HashAlgoSha1 {
return hash.Value
}
}
return ""
}

// GetMd5Hash returns the MD5 hash from a file's hashes if available
func GetMd5Hash(file File) string {
for _, hash := range file.Hashes {
if hash.Algo == HashAlgoMd5 {
return hash.Value
}
}
return ""
}

// HasGameVersion checks if a file supports a specific game version
func (f File) HasGameVersion(version string) bool {
for _, v := range f.GameVersions {
if v == version {
return true
}
}
return false
}

// HasModLoader checks if a file supports a specific mod loader
func (f File) HasModLoader(loader ModLoaderType) bool {
// Check GameVersions for loader name
loaderName := loader.String()
for _, v := range f.GameVersions {
if v == loaderName {
return true
}
}
return false
}
