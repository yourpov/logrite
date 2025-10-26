package logrite

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

// ── Configuration ─────────────────────────────────────

type Config struct {
	ShowIcons    bool
	UppercaseTag bool
	UseColors    bool
}

var config = Config{
	ShowIcons:    true,
	UppercaseTag: true,
	UseColors:    true,
}

func SetConfig(c Config) {
	config = c
}

// ── Named Color Aliases ───────────────────────────────

// Foreground
var (
	Magenta      = color.FgHiMagenta
	Red          = color.FgRed
	Green        = color.FgGreen
	Yellow       = color.FgYellow
	Cyan         = color.FgCyan
	White        = color.FgWhite
	Black        = color.FgBlack
	BrightRed    = color.FgHiRed
	BrightCyan   = color.FgHiCyan
	BrightGreen  = color.FgHiGreen
	BrightYellow = color.FgHiYellow
	BrightWhite  = color.FgHiWhite
)

// Background
var (
	BgPink         = color.BgHiMagenta
	BgRed          = color.BgRed
	BgGreen        = color.BgGreen
	BgYellow       = color.BgYellow
	BgCyan         = color.BgCyan
	BgWhite        = color.BgWhite
	BgBlack        = color.BgBlack
	BgBrightRed    = color.BgHiRed
	BgBrightGreen  = color.BgHiGreen
	BgBrightYellow = color.BgHiYellow
	BgBrightCyan   = color.BgHiCyan
)

// ── Internal Formatter ─────────────────────────────────

func logLine(tag, message string, bg, fg color.Attribute) {
	normalizedTag := strings.ToLower(tag)

	icons := map[string]string{
		"info": "ℹ️ ", "warn": "⚠️ ", "error": "❌", "success": "✅",
		"login": "🔐", "web": "🌐", "command": "💻", "debug": "🐞",
	}

	icon := ""
	if config.ShowIcons {
		if val, ok := icons[normalizedTag]; ok {
			icon = val + " "
		}
	}

	displayTag := tag
	if config.UppercaseTag {
		displayTag = strings.ToUpper(tag)
	}

	label := fmt.Sprintf(" %-8s", displayTag)

	if config.UseColors {
		fmt.Println(icon+color.New(bg, fg).Sprint(label), message)
	} else {
		fmt.Println(icon+label, message)
	}
}

// ── Custom Tagged Log ─────────────────────────────────

func Log(tag, message string, bg, fg color.Attribute) {
	logLine(tag, message, bg, fg)
}

func Custom(emoji, tag, format string, fg, bg color.Attribute, args ...any) {
	displayTag := tag
	if config.UppercaseTag {
		displayTag = strings.ToUpper(tag)
	}

	label := fmt.Sprintf(" %-8s", displayTag)
	message := fmt.Sprintf(format, args...)

	if config.UseColors {
		fmt.Println(emoji+" "+color.New(bg, fg).Sprint(label), message)
	} else {
		fmt.Println(emoji+" "+label, message)
	}
}

// ── Generic Shortcuts ───────────────────────────

func Info(msg string, args ...interface{}) {
	logLine("info", fmt.Sprintf(msg, args...), BgWhite, Black)
}

func Warn(msg string, args ...interface{}) {
	logLine("warn", fmt.Sprintf(msg, args...), BgYellow, Black)
}

func Error(msg string, args ...interface{}) error {
	logLine("error", fmt.Sprintf(msg, args...), BgBrightRed, White)
	return fmt.Errorf(msg, args...)
}

func Success(msg string, args ...interface{}) {
	logLine("success", fmt.Sprintf(msg, args...), BgGreen, Black)
}

func Debug(msg string, args ...interface{}) {
	logLine("debug", fmt.Sprintf(msg, args...), BgPink, White)
}


// ── Example Custom Functions Below ────────────────────────────────────────

// Command Executions
func Cmd(path, duration, user string) {
	logLine("command",
		fmt.Sprintf("'%s' executed in %s by %s",
			color.New(BrightGreen).Sprint(path),
			color.New(BrightGreen).Sprint(duration),
			color.New(BrightGreen).Sprint(user),
		),
		BgBrightGreen, Black,
	)
}

// Login Events
func Login(user, ip string) {
	logLine("login",
		fmt.Sprintf("Member %s logged in from %s",
			color.New(BrightYellow).Sprint(user),
			color.New(BrightYellow).Sprint(ip),
		),
		BgBrightYellow, Black,
	)
}

// Webserver Startup
func WebStart(port int) {
	logLine("web",
		fmt.Sprintf("Started webserver on port %s",
			color.New(BrightCyan).Sprint(port)),
		BgCyan, White,
	)
}
