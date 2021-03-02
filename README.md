# Just Notes On Comptuer (JNOC)

## Overview

This is a very simple implementation of the Zettlekasten, Slip Box, and Smart Notes that's on your local machine. 

Advantages

* You don't have to worry about your note taking apps becoming outdated. Text files with Markdown are here to stay.
* You can easily write your own scripts to interact with the files
* You can edit notes with the editor of your choosing (VS Code, vim, emacs, Atom, Typora, etc.)

Simply clone this repo and begin creating your personal notes system. The repo inclusdes the `jnoc` binary that helps you create the notes. 

For more information run `./jnoc --help`

## Navigating Between Notes
### VS Code
`cmd+click` on a path to a note to open it. 
### Vim
Move the cursor over the path and type `gf`.

## Note Structure

* Title
* Metadata
* Content
* References

Example:
```
# Negative Visualizations  
---
**Metadata**  
ID: 202004291359  
Tags: [ #NegativeVisualization #Tranquility ]  
Related Notes:  
[link 1](../AGuideToTheGoodLife-202004291355.md)  

---

Negative Visualization is a technique...

---

[AGuideToTheGoodLife]: . "William B. Irvine *A Guide to the Good Life: The Ancient Art of Stoic Joy*)"  
```

## Usage

### Creating a New Note
```
./jnoc note
Note Topic: Curriculum
Note: Zettles/Curriculum-20200505114433.md
```

### Note Structure
