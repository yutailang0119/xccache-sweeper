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

func cache_paths() ([]string, error) {
  xcode_build_location_sytle, err := exec.Command("defaults", "read", "com.apple.dt.Xcode", "IDEBuildLocationStyle").Output()
  if err != nil {
    return nil, err
  }

  usr, _ := user.Current()
  if string(xcode_build_location_sytle) == "Unique" {
    return []string{strings.Replace("~/Library/Developer/Xcode",  "~", usr.HomeDir, 1)}, nil
  } else {
    // TODO: 再帰的に*.xcodeprojを探す

    paths []string := []
    err = filepath.Walk(usr.HomeDir, func(path string, info os.FileInfo, err error) error {
      if info.IsDir()  {

        ok, err := filepath.Match("*.xcodeproj", path)
        if err != nil {
          return err
        }

        if ok {
          err := os.RemoveAll(path)
          if err != nil {
            return err
          }
          log.Println(path)
          return filepath.SkipDir
        }
      
      }
      
      return nil
  })

  if err != nil {
    return err
  }

    usr, _ := user.Current()
    return []string{strings.Replace("~/Library/Developer/Xcode",  "~", usr.HomeDir, 1)}, nil
  }
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

  xcode_build_location_sytle, err := exec.Command("defaults", "read", "com.apple.dt.Xcode", "IDEBuildLocationStyle").Output()
  if err != nil {
    panic(err)
  }

  var xcode_caches_path string
  if string(xcode_build_location_sytle) == "Unique" {
    usr, _ := user.Current()
    xcode_caches_path = strings.Replace("~/Library/Developer/Xcode",  "~", usr.HomeDir, 1)
  } else {
    usr, _ := user.Current()
    xcode_caches_path = strings.Replace("~/Library/Developer/Xcode",  "~", usr.HomeDir, 1)
  }

  derived_data_path := filepath.Join(xcode_caches_path, "DerivedData")
  archives_path := filepath.Join(xcode_caches_path, "Archives")
  matching_archives_path := filepath.Join(archives_path, "*")

  now := time.Now()
  expired := now.AddDate(0, -1, 0)

  err = filepath.Walk(xcode_caches_path, func(path string, info os.FileInfo, err error) error {
    if info.IsDir()  {

      ok, err := filepath.Match(derived_data_path, path)
      if err != nil {
        return err
      }

      if ok {
        err := os.RemoveAll(path)
        if err != nil {
          return err
        }
        log.Println(path)
        return filepath.SkipDir
      }

      ok, err = filepath.Match(matching_archives_path, path)
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

