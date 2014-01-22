# Keigo -- [Keiko-chan] client library

# Intro
keigo is [keiko-chan] CLI client written in go.

# Install

```
go get github.com/taichi/keigo
go install github.com/taichi/keigo
```

# Usage

## Set Up
make your **config.toml**.
The config file uses [TOML](https://github.com/mojombo/toml) syntax. 

like below

```toml
keiko_address="keiko.example.com:60000"
```

## Configurable Options

* keiko_address
    * **required**
    * host name and port no of keiko-chan
* terminal_code
    * default value is \\r(13)
* timeout
    * default value is 60 seconds
    * waiting specific Millisecond for response from the [keiko-chan]
* retry
    * default value is 3 times
    * retry specific times

## Commands
keigo depends on [spf13/cobra](https://github.com/spf13/cobra).

keigo supports following sub commands.

## high level sub commands

### help
Print command usage and options.

```
> keigo help red
turn on/off red lamp.
Example:
    keigo red off
    keigo red on
    keigo red blink
    keigo red quickblink
    keigo red on -t 3 -w 2

Usage:
  keigo red [flags]

 Available Flags:
  -c, --config="config.toml": config file path
  -t, --time=0: turn lamp off after specific seconds.
  -v, --verbose=false: print debug informations
  -w, --wait=0: waiting for specific seconds.


Use "keigo help [command]" for more information about that command.
```

### version
Print the Version No of Keigo

```
> keigo version
Keigo     : 0.0.1
```
    

### info
Print the information of Keiko-chan

```
> keigo info
Version             : 14.220.1G
UNIT ID             : 1510
Effective Date      : Not registered
Contract Number     : Not registered
Model Name          : DN-1510GL
Production          : 1301
Serial Number       : 0000000000
```

###  red/green/yellow/buzzer1/buzzer2
turn on/off lamp or buzzer.
there commands supports 

* off
* on
* blink
* quickblink


```
> keigo red on
Current Lamp Status
-------------------
RED        :    on
-------------------
YELLOW     :    off
-------------------
GREEN      :    off
-------------------
BUZZER1    :    off
-------------------
BUZZER2    :    off

>keigo yellow blink
Current Lamp Status
-------------------
RED        :    off
-------------------
YELLOW     :  blink
-------------------
GREEN      :    off
-------------------
BUZZER1    :    off
-------------------
BUZZER2    :    off
```

### off
turn off all lamps and buzzers.

```
> keigo off
Current Lamp Status
-------------------
RED        :    off
-------------------
YELLOW     :    off
-------------------
GREEN      :    off
-------------------
BUZZER1    :    off
-------------------
BUZZER2    :    off
```

### play
play configured sound.

```
> keigo play 1 -t 1
State      : playing
SoundNo    : 01
Repeat?    : false
PlayTimes  : 01
```

### stop
stop sound.

```
> keigo stop
State      : stopped
SoundNo    : 01
Repeat?    : false
PlayTimes  : 01
```

### low level sub commands
you should see the [official manual(pdf)](http://www.isa-j.co.jp/dn1510gl/files/dn1510gl-manual-20130426.pdf) and [codes](https://github.com/keigo/command).



[keiko-chan]: http://www.isa-j.co.jp/product/keiko/