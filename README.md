# go-xcodecache-sweeper

# How to use

## Builde main file

```bash
$go build main.go
```

## Execute app

```bash
$./main
```

## Use with Automator (Recommend)

1. Launch to Automator
2. Choose a type of Application  
    ![AutomatorApplication.png](./Resources/Automator/AutomatorApplication.png)
3. Select `Run Shell Script`
4. Wirte a Shell command
    ```bash
    ~/go/src/github.com/yutailang0119/go-xcodecache-sweeper/main // WorkingDirectory/go-xcodecache-sweeper/main
    ```
    ![ShellCommand.png](./Resources/Automator/ShellCommand.png)
5. Save as Application  
    ![SaveAsApplication.png](./Resources/Automator/SaveAsApplication.png)
6. `System Preference > Users & Groups > Login Items`  
    ![SelectApplication.png](./Resources/Automator/SelectApplication.png)
7. Select this app

