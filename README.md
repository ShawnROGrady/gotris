```
   ______    ______     _     
  / ____/___/_  __/____(_)____
 / / __/ __ \/ / / ___/ / ___/
/ /_/ / /_/ / / / /  / (__  ) 
\____/\____/_/ /_/  /_/____/  
```

# Overview
A terminal based version of the 1984 game written in Go

## Compatibility 
This has been primarily tested on macOS Mojave with `$TERM=xterm-256color`. While I have been able to test on an Ubuntu VM I am sure syscalls and keyboard inputs vary depending on `$TERM`, which I haven't accounted for.

Windows is not currently supported.

## Demo
![demo-gif](https://github.com/ShawnROGrady/gotris/blob/add-readme/assets/gotris-demo.gif)
# Configuration
## Options
1. `-disable-ghost`: Don't show the 'ghost' of the current piece
2. `-disable-side`: Don't show the side bar (next piece, current score, and controls)
3. `-light-mode`: Update colors to work for light color schemes
4. `-low-contrast`: Update colors to use lower contrast (updates background to white for 'light-mode', black otherwise)
5. `-scheme`: The control scheme to use, multiple may be specified (default: home-row)
   -  all schemes can be viewed using `-describe-scheme` sub-command described below

Quick comparison of the color options (`-light-mode` enabled on bottom, `-low-contrast` enabled on right):
![colors](https://github.com/ShawnROGrady/gotris/blob/add-readme/assets/gotris-colors.png)

**NOTE:** you will also see `-debug` and `-cpuprofile` listed after running `gotris -h`. These were for my personal use when building this and won't be of much use to the standard user.
## Sub-commands
1. `-colors`: Display the colors that will be used throughout the game then exit
   - useful for previewing results of the `-disable-ghost`, `-disable-side`, `-light-mode`, and `-low-contrast` options
2. `-describe-scheme`: Prints the specified control scheme then exits. If none specified then all available schemes are described

# Some Notes
The goal of this project was to create a fully featured yet simple terminal-based version of tetris written in Go using just the standard library. This project was meant to be a learning experience which is why I chose to avoid external depencies.

## Planned Features
Easier:
- [ ] Initial difficulty selection
- [ ] Width+height selection
- [ ] Other display options (opacity of ghost piece, monochrome mode, etc.)

Harder:
- [ ] Windows support
- [ ] "Hold" piece option
- [ ] Pause+resume
