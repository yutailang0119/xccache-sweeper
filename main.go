package main

import (
  "os"
  "os/user"
  "os/exec"
  "log"
  "strings"
  "strconv"
  "time"
  "path/filepath"
)

func cached_derived_data_paths() ([]string, error) {
  xcode_build_location_style, err := exec.Command("defaults", "read", "com.apple.dt.Xcode", "IDEBuildLocationStyle").Output()
  if err != nil {
    return nil, err
  }

  usr, _ := user.Current()
  if strings.TrimSpace(string(xcode_build_location_style)) == "Unique" {
    return []string{strings.Replace("~/Library/Developer/Xcode/DerivedData",  "~", usr.HomeDir, 1)}, nil
  } else {
    paths := []string{}
    err := filepath.Walk(usr.HomeDir, func(path string, info os.FileInfo, err error) error {
      if info.IsDir()  {

        if filepath.Ext(path) == ".xcodeproj" {
          cmd := "xcodebuild -project " + path
          cmd = cmd + " -showBuildSettings | grep -e \"BUILD_ROOT\""
          build_root, _:= exec.Command("sh", "-c", cmd).Output()
          output := strings.TrimSpace(string(build_root))

          build_root_path := strings.TrimPrefix(output, "BUILD_ROOT = ")
          derived_data_path := strings.TrimSuffix(build_root_path, "/Build/Products")

          paths = append(paths, derived_data_path)

          return filepath.SkipDir
        }

      }

      return nil
    })

    if err != nil {
      return nil, err
    }

    return paths, nil

  }
}

func cached_archives_path() string {

  usr, _ := user.Current()
  xcode_plist_path := strings.Replace("~/Library/Preferences/com.apple.dt.Xcode",  "~", usr.HomeDir, 1)
  archives_path, err := exec.Command("defaults", "read", xcode_plist_path, "IDECustomDistributionArchivesLocation").Output()
  if err != nil {
    return strings.Replace("~/Library/Developer/Xcode/Archives",  "~", usr.HomeDir, 1)
  }
  return string(archives_path)

}

func check_expired(dir string, expired time.Time) (bool, error) {

  splited := strings.Split(dir, "-")
  year, err := strconv.Atoi(splited[0])
  if err != nil {
    return false, err
  }

  month, err := strconv.Atoi(splited[1])
  if err != nil {
    return false, err
  }

  day, err := strconv.Atoi(splited[2])
  if err != nil {
    return false, err
  }

  date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

  return date.Before(expired), nil
}

func main() {

  derived_data_paths, err := cached_derived_data_paths()
  if err != nil {
    panic(err)
  }

  for _, path := range derived_data_paths {
    err := os.RemoveAll(path)
    if err == nil {
      log.Println(path)
    }
  }

  archives_path := cached_archives_path()
  matching_archives_path := filepath.Join(archives_path, "*")

  now := time.Now()
  expired := now.AddDate(0, -1, 0)

  err = filepath.Walk(archives_path, func(path string, info os.FileInfo, err error) error {
    if info.IsDir()  {

      ok, err := filepath.Match(matching_archives_path, path)
      if err != nil {
        return err
      }

      if ok {
        dir := strings.Replace(path, archives_path + "/", "", 1)
        is_expired, err := check_expired(dir, expired)
        if err != nil {
          return err
        }

        if is_expired {
          err := os.RemoveAll(path)
          if err != nil {
            return err
          }
          log.Println(path)
          return filepath.SkipDir
        }
      }
    }

    return nil
  })

  if err != nil {
    panic(err)
  }

}

