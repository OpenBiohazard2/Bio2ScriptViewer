# Resident Evil 2 Script Viewer

<div style="display:inline-block;">
<img src="https://github.com/samuelyuan/Bio2ScriptViewer/raw/master/screenshots/ScriptViewer.png" alt="ScriptViewer" width="400" height="300" />
</div>

## About

You can view the script files in the original Resident Evil 2 / Biohazard 2 as pseudocode next to the original bytecode. Every function is placed in a separate file to make it easier to switch between functions.

The script data is stored as part of the room description file (.RDT). When you open any RDT file, the script files will be extracted from the RDT and the list of files is shown on the left.

## Scripting Engine

This script viewer will make it easier for anyone to understand the scripting logic used by the original Resident Evil 2 game. 

The script in the RDT file is originally stored as one block of hex values, but the ScriptViewer parses the data by splitting it into script commands, and combining script commands into functions using the EvtEnd opcode (0x01) as the separator.

* "init.scd" is the script that executes when the player initially enters a room and only executes once. This script is used for static events.

* "sub0.scd" and "sub1.scd" are scripts that will start running while the player is in the room as part of the main game loop. The script engine creates two events and runs each event separately on its own ScriptThread until the function exits, i.e. sub0.scd will run on ScriptThread0 and sub1.scd will run on ScriptThread1. These scripts can either spawn other events, which starts a separate ScriptThread, or call other functions on the same ScriptThread. The sub scripts are used for dynamic events.

The left panel shows the original bytecode in hexadecimal, which is the same data that can be found if you open the RDT file in an hex editor and search for the sequence of bytes. 

The right panel shows the corresponding pseudocode that contains a function name and its parameters. The first hex value in each row is the opcode and the subsequent hex values after the opcode are the function parameters. The opcode parameters are determined in advance by the scripting engine, and the parameter types can be 8 bit, 16 bit, or 32 bit values.

