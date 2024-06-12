# ASCII-ART PROGRAM

The program receives a string as an argument and outputs the string in a graphic representation using ASCII.

## Setup Instructions

1. Clone the repository:
   ```sh
   git clone https://learn.zone01kisumu.ke/git/fgitonga/ascii-art-fs.git
   ```
   
2. Navigate to the project directory:
   ```sh
   cd ascii-art-fs
   ```
   
3. Install dependencies:
   ```sh
   go mod tidy
   ```

## Usage

To run a go project locally, execute:
  ```sh
  go run main.go [STRING][BANNER]
  ```
EX:

```
$ go run . "hello" standard | cat -e
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $
$ go run . "Hello There!" shadow | cat -e
                                                                                         $
_|    _|          _| _|                _|_|_|_|_| _|                                  _| $
_|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   _| $
_|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_| _| $
_|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|          $
_|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_| _| $
                                                                                         $
                                                                                         $
$ go run . "Hello There!" thinkertoy | cat -e
                                                $
o  o     o o           o-O-o o                o $
|  |     | |             |   |                | $
O--O o-o | | o-o         |   O--o o-o o-o o-o o $
|  | |-' | | | |         |   |  | |-' |   |-'   $
o  o o-o o o o-o         o   o  o o-o o   o-o O $
                                                $
                                                $
```

## Fs
```
Ascii-art/
|----main.go
|----standard.txt
|----shadow.txt
|----thinkertoy.txt
|----utils
|   |----func.go
|   |----test_func.go
|----README.md
|----go.sum
|____go.mod
```
## Contributors
1. Godwin Ouma
2. Valeria Muhembele
3. Fred Gitonga
