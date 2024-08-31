package utils

import (
    "time"
    "fmt"
)

// 現在のUTC時間を返す
func GetCurrentUTCTime() time.Time {
    return time.Now().UTC()
}

// 現在のローカル時間を返す
func GetCurrentLocalTime() time.Time {
    return time.Now()
}

// 指定されたフォーマットで現在の時間を文字列で返す
func GetFormattedCurrentTime(format string) string {
    return time.Now().Format(format)
}

// タイムスタンプを指定されたフォーマットで文字列に変換
func FormatTimestamp(timestamp time.Time, format string) string {
    return timestamp.Format(format)
}

// 文字列を指定されたフォーマットでパースしてタイムスタンプに変換
func ParseTimestamp(timestampStr string, format string) (time.Time, error) {
    return time.Parse(format, timestampStr)
}

// 2つのタイムスタンプの差分を返す
func GetTimeDifference(t1, t2 time.Time) time.Duration {
    return t1.Sub(t2)
}

// タイムスタンプに指定された時間を追加
func AddDurationToTimestamp(timestamp time.Time, duration time.Duration) time.Time {
    return timestamp.Add(duration)
}

// ISO8601形式の現在の時間を返す
func GetCurrentISO8601Time() string {
    return time.Now().UTC().Format(time.RFC3339)
}

// 指定されたタイムゾーンの現在の時間を返す
func GetCurrentTimeInTimeZone(timeZone string) (time.Time, error) {
    loc, err := time.LoadLocation(timeZone)
    if err != nil {
        return time.Time{}, err
    }
    return time.Now().In(loc), nil
}

// 日付の開始時間を返す（つまり、00:00:00）
func GetStartOfDay(t time.Time) time.Time {
    year, month, day := t.Date()
    return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

// 日付の終了時間を返す（つまり、23:59:59）
func GetEndOfDay(t time.Time) time.Time {
    year, month, day := t.Date()
    return time.Date(year, month, day, 23, 59, 59, 999999999, t.Location())
}

// 日付が今日かどうかを確認
func IsToday(t time.Time) bool {
    now := time.Now()
    return t.Year() == now.Year() && t.YearDay() == now.YearDay()
}

// UNIXタイムスタンプを時間に変換
func UnixToTime(unixTimestamp int64) time.Time {
    return time.Unix(unixTimestamp, 0)
}

// 時間をUNIXタイムスタンプに変換
func TimeToUnix(t time.Time) int64 {
    return t.Unix()
}

// 月初を取得
func GetStartOfMonth(t time.Time) time.Time {
    year, month, _ := t.Date()
    return time.Date(year, month, 1, 0, 0, 0, 0, t.Location())
}

// 月末を取得
func GetEndOfMonth(t time.Time) time.Time {
    year, month, _ := t.Date()
    firstDayOfNextMonth := time.Date(year, month+1, 1, 0, 0, 0, 0, t.Location())
    return firstDayOfNextMonth.Add(-time.Nanosecond)
}

// 年初を取得
func GetStartOfYear(t time.Time) time.Time {
    year := t.Year()
    return time.Date(year, 1, 1, 0, 0, 0, 0, t.Location())
}

// 年末を取得
func GetEndOfYear(t time.Time) time.Time {
    year := t.Year()
    return time.Date(year, 12, 31, 23, 59, 59, 999999999, t.Location())
}

// タイムゾーンを変換
func ConvertTimeToTimeZone(t time.Time, timeZone string) (time.Time, error) {
    loc, err := time.LoadLocation(timeZone)
    if err != nil {
        return time.Time{}, err
    }
    return t.In(loc), nil
}

// 日付をフォーマットして表示
func FormatDate(t time.Time) string {
    return t.Format("2006-01-02")
}

// 時間をフォーマットして表示
func FormatTime(t time.Time) string {
    return t.Format("15:04:05")
}

// 日付と時間をフォーマットして表示
func FormatDateTime(t time.Time) string {
    return t.Format("2006-01-02 15:04:05")
}

// 指定された日数を追加
func AddDays(t time.Time, days int) time.Time {
    return t.AddDate(0, 0, days)
}

// 指定された月数を追加
func AddMonths(t time.Time, months int) time.Time {
    return t.AddDate(0, months, 0)
}

// 指定された年数を追加
func AddYears(t time.Time, years int) time.Time {
    return t.AddDate(years, 0, 0)
}

// 現在の時間をフォーマットして表示
func PrintCurrentTime() {
    fmt.Println("Current Time:", time.Now().Format("2006-01-02 15:04:05"))
}
