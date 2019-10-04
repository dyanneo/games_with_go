#Games with Go

https://www.youtube.com/watch?v=OXSMx45kayw&list=PLDZujg-VgQlZUy1iCqBbe5faZLMkA3g2x&index=8&t=0s

^ for episode 6; around 6 min mark note the mingw setup for windows

make sure to execute `go get github.com/veandco/go-sdl2/sdl`

### 10/4/19 plan:  

*OS abstraction layer*

* draw pixels to screen  
* play sounds  
* get input (keyboards/mouse/controllers/touch screen)  
* read from files  
* networking  
* threading  



*Game Loops*

Episode 7
```
for {
    Update
        get all inputs (keyboard, controllers, etc)
        update all the things like physics and ai
    Draw
        draw everything
}
```