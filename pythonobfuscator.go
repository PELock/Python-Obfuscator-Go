/******************************************************************************
 * Python Obfuscator WebApi interface
 *
 * Version        : v1.0.0
 * Language       : Go
 * Author         : Bartosz Wójcik
 * Web page       : https://www.pelock.com
 *
 *****************************************************************************/

package pythonobfuscator

import (
	"bytes"
	"compress/zlib"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	APIURL = "https://www.pelock.com/api/python-obfuscator/v1"

	ErrorSuccess     = 0
	ErrorInputSize   = 1
	ErrorInput       = 2
	ErrorParsing     = 3
	ErrorObfuscation = 4
	ErrorOutput      = 5

	CodeVirtualizationVM   = "vm"
	CodeVirtualizationFSA  = "fsa"
	CodeVirtualizationFLAT = "flat"

	RenameStyleIL         = "il"
	RenameStyleO0         = "o0"
	RenameStyleConfusable = "confusable"
	RenameStyleHex        = "hex"
	RenameStyleHomoglyph  = "homoglyph"
	RenameStyleMangled    = "mangled"
)

// Result is the parsed Web API JSON response.
type Result struct {
	Error             int    `json:"error"`
	Output            string `json:"output,omitempty"`
	Demo              bool   `json:"demo,omitempty"`
	LicenseExpiration string `json:"license_expiration,omitempty"`
	UsagesTotal       int    `json:"usages_total,omitempty"`
	StringLimit       int    `json:"string_limit,omitempty"`
}

// Client is the Python Obfuscator Web API client.
// Defaults match the JavaScript SDK.
type Client struct {
	APIKey     string
	APIURL     string
	UserAgent  string
	HTTPClient *http.Client

	EnableCompression bool

	// Seed is a fixed random seed for reproducible output. Nil means unset.
	Seed *int
	// RandomizationDensity is intensity 0-100. Nil means unset.
	RandomizationDensity *int
	RenameStyle          string
	CodeVirtualization   string

	SelfDefending    bool
	ProtectionLinker bool

	RenameVariables     bool
	RenameParameters    bool
	RenameFunctions     bool
	RenameFunctionCalls bool
	ShuffleFunctions    bool
	ResolveConstants    bool

	SplitStrings          bool
	ModifyStrings         bool
	EncryptStrings        bool
	StringCharArrayVault  bool

	EncryptIntegers     bool
	EncryptFloating     bool
	MBABinops           bool
	IntegersToFloating  bool
	IntegersToArrays    bool
	FloatsToArrays      bool
	AffineIntegerMask   bool

	ArrayIntCrypt    bool
	ArrayCharCrypt   bool
	ArrayDoubleCrypt bool
	ArrayStringCrypt bool

	InsertRandomValueBucket  bool
	RandomBucketIntegers     bool
	RandomBucketArrays       bool
	RandomBucketFunctions    bool
	RandomBucketCharacters   bool
	RandomBucketAntiRegex    bool
	RandomBucketAutostart    bool

	InsertTernaryOperators bool
	ComplexifyBooleans     bool
	OpaqueBranches         bool
	OpaqueMixerChain       bool
	InsertDeadCode         bool
	TryFinallyNoise        bool

	DecoyFunctions            bool
	LambdaDecoys              bool
	LiteralPadding            bool
	FakeImportMarkers         bool
	DynamicGetattrCalls       bool
	ObfuscateImports          bool
	CallbackRegistrationStubs bool

	DetectDebugger bool
	AntiVM         bool
	AntiSandbox    bool
	AntiEmulator   bool

	RemoveComments bool
}

// New creates a client. An empty or invalid key runs demo mode.
func New(apiKey string) *Client {
	return &Client{
		APIKey:                    apiKey,
		APIURL:                    APIURL,
		UserAgent:                 "PELock Python Obfuscator",
		HTTPClient:                &http.Client{Timeout: 120 * time.Second},
		EnableCompression:         false,
		CodeVirtualization:        CodeVirtualizationVM,
		RenameVariables:           true,
		RenameParameters:          true,
		RenameFunctions:           true,
		RenameFunctionCalls:       true,
		ShuffleFunctions:          true,
		ResolveConstants:          true,
		SplitStrings:              true,
		ModifyStrings:             true,
		EncryptStrings:            true,
		StringCharArrayVault:      true,
		EncryptIntegers:           true,
		EncryptFloating:           true,
		MBABinops:                 true,
		IntegersToFloating:        true,
		IntegersToArrays:          true,
		FloatsToArrays:            true,
		AffineIntegerMask:         true,
		ArrayIntCrypt:             true,
		ArrayCharCrypt:            true,
		ArrayDoubleCrypt:          true,
		ArrayStringCrypt:          true,
		InsertRandomValueBucket:   true,
		RandomBucketIntegers:      true,
		RandomBucketArrays:        true,
		RandomBucketFunctions:     true,
		RandomBucketCharacters:    true,
		RandomBucketAntiRegex:     true,
		RandomBucketAutostart:     true,
		InsertTernaryOperators:    true,
		ComplexifyBooleans:        true,
		OpaqueBranches:            true,
		OpaqueMixerChain:          true,
		InsertDeadCode:            true,
		TryFinallyNoise:           true,
		DecoyFunctions:            true,
		LambdaDecoys:              true,
		LiteralPadding:            true,
		FakeImportMarkers:         true,
		DynamicGetattrCalls:       true,
		ObfuscateImports:          true,
		CallbackRegistrationStubs: true,
		RemoveComments:            true,
	}
}

// Login returns license and quota information for the activation key.
func (c *Client) Login(ctx context.Context) (*Result, error) {
	return c.postRequest(ctx, map[string]string{"command": "login"})
}

// ObfuscateScriptFile reads a UTF-8 Python script and obfuscates it.
func (c *Client) ObfuscateScriptFile(ctx context.Context, scriptFilePath string) (*Result, error) {
	source, err := os.ReadFile(scriptFilePath)
	if err != nil {
		return nil, err
	}
	if len(source) == 0 {
		return nil, fmt.Errorf("empty source file")
	}
	return c.ObfuscateScriptSource(ctx, string(source))
}

// ObfuscateScriptSource obfuscates Python source code.
func (c *Client) ObfuscateScriptSource(ctx context.Context, scriptSource string) (*Result, error) {
	return c.postRequest(ctx, map[string]string{
		"command": "obfuscate",
		"source":  scriptSource,
	})
}

func (c *Client) postRequest(ctx context.Context, params map[string]string) (*Result, error) {
	if c.APIKey != "" {
		params["key"] = c.APIKey
	}

	if c.Seed != nil {
		params["seed"] = strconv.Itoa(*c.Seed)
	}
	if c.RandomizationDensity != nil {
		params["randomization_density"] = strconv.Itoa(*c.RandomizationDensity)
	}
	if c.RenameStyle != "" {
		params["rename_style"] = c.RenameStyle
	}
	if c.CodeVirtualization != "" {
		params["code_virtualization"] = c.CodeVirtualization
	}

	if c.SelfDefending {
		params["self_defending"] = "1"
	}
	if c.ProtectionLinker {
		params["protection_linker"] = "1"
	}
	if c.RenameVariables {
		params["rename_variables"] = "1"
	}
	if c.RenameParameters {
		params["rename_parameters"] = "1"
	}
	if c.RenameFunctions {
		params["rename_functions"] = "1"
	}
	if c.RenameFunctionCalls {
		params["rename_function_calls"] = "1"
	}
	if c.ShuffleFunctions {
		params["shuffle_functions"] = "1"
	}
	if c.ResolveConstants {
		params["resolve_constants"] = "1"
	}
	if c.SplitStrings {
		params["split_strings"] = "1"
	}
	if c.ModifyStrings {
		params["modify_strings"] = "1"
	}
	if c.EncryptStrings {
		params["encrypt_strings"] = "1"
	}
	if c.StringCharArrayVault {
		params["string_char_array_vault"] = "1"
	}
	if c.EncryptIntegers {
		params["encrypt_integers"] = "1"
	}
	if c.EncryptFloating {
		params["encrypt_floating"] = "1"
	}
	if c.MBABinops {
		params["mba_binops"] = "1"
	}
	if c.IntegersToFloating {
		params["integers_to_floating"] = "1"
	}
	if c.IntegersToArrays {
		params["integers_to_arrays"] = "1"
	}
	if c.FloatsToArrays {
		params["floats_to_arrays"] = "1"
	}
	if c.AffineIntegerMask {
		params["affine_integer_mask"] = "1"
	}
	if c.ArrayIntCrypt {
		params["array_int_crypt"] = "1"
	}
	if c.ArrayCharCrypt {
		params["array_char_crypt"] = "1"
	}
	if c.ArrayDoubleCrypt {
		params["array_double_crypt"] = "1"
	}
	if c.ArrayStringCrypt {
		params["array_string_crypt"] = "1"
	}
	if c.InsertRandomValueBucket {
		params["insert_random_value_bucket"] = "1"
	}
	if c.RandomBucketIntegers {
		params["random_bucket_integers"] = "1"
	}
	if c.RandomBucketArrays {
		params["random_bucket_arrays"] = "1"
	}
	if c.RandomBucketFunctions {
		params["random_bucket_functions"] = "1"
	}
	if c.RandomBucketCharacters {
		params["random_bucket_characters"] = "1"
	}
	if c.RandomBucketAntiRegex {
		params["random_bucket_anti_regex"] = "1"
	}
	if c.RandomBucketAutostart {
		params["random_bucket_autostart"] = "1"
	}
	if c.InsertTernaryOperators {
		params["insert_ternary_operators"] = "1"
	}
	if c.ComplexifyBooleans {
		params["complexify_booleans"] = "1"
	}
	if c.OpaqueBranches {
		params["opaque_branches"] = "1"
	}
	if c.OpaqueMixerChain {
		params["opaque_mixer_chain"] = "1"
	}
	if c.InsertDeadCode {
		params["insert_dead_code"] = "1"
	}
	if c.TryFinallyNoise {
		params["try_finally_noise"] = "1"
	}
	if c.DecoyFunctions {
		params["decoy_functions"] = "1"
	}
	if c.LambdaDecoys {
		params["lambda_decoys"] = "1"
	}
	if c.LiteralPadding {
		params["literal_padding"] = "1"
	}
	if c.FakeImportMarkers {
		params["fake_import_markers"] = "1"
	}
	if c.DynamicGetattrCalls {
		params["dynamic_getattr_calls"] = "1"
	}
	if c.ObfuscateImports {
		params["obfuscate_imports"] = "1"
	}
	if c.CallbackRegistrationStubs {
		params["callback_registration_stubs"] = "1"
	}
	if c.DetectDebugger {
		params["detect_debugger"] = "1"
	}
	if c.AntiVM {
		params["anti_vm"] = "1"
	}
	if c.AntiSandbox {
		params["anti_sandbox"] = "1"
	}
	if c.AntiEmulator {
		params["anti_emulator"] = "1"
	}
	if c.RemoveComments {
		params["remove_comments"] = "1"
	}

	if c.EnableCompression && params["source"] != "" {
		compressed, err := zlibCompressBase64(params["source"])
		if err != nil {
			return nil, err
		}
		params["source"] = compressed
		params["compression"] = "1"
	}

	body, err := c.postMultipart(ctx, params)
	if err != nil {
		return nil, err
	}
	if len(body) == 0 {
		return nil, fmt.Errorf("empty API response")
	}

	var result Result
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if c.EnableCompression && result.Error == ErrorSuccess && result.Output != "" {
		plain, err := zlibDecompressBase64(result.Output)
		if err != nil {
			return nil, err
		}
		result.Output = plain
	}

	return &result, nil
}

func (c *Client) postMultipart(ctx context.Context, fields map[string]string) ([]byte, error) {
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	for k, v := range fields {
		if err := w.WriteField(k, v); err != nil {
			return nil, err
		}
	}
	if err := w.Close(); err != nil {
		return nil, err
	}

	url := c.APIURL
	if url == "" {
		url = APIURL
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, &buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	httpClient := c.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func zlibCompressBase64(s string) (string, error) {
	var buf bytes.Buffer
	zw, err := zlib.NewWriterLevel(&buf, zlib.BestCompression)
	if err != nil {
		return "", err
	}
	if _, err := zw.Write([]byte(s)); err != nil {
		_ = zw.Close()
		return "", err
	}
	if err := zw.Close(); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

func zlibDecompressBase64(s string) (string, error) {
	raw, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	zr, err := zlib.NewReader(bytes.NewReader(raw))
	if err != nil {
		return "", err
	}
	defer zr.Close()
	plain, err := io.ReadAll(zr)
	if err != nil {
		return "", err
	}
	return string(plain), nil
}
