$(document).ready(function(){

    var canvas = $("#canvas")[0];
    var ctx = canvas.getContext("2d");
    var width = $("#canvas").width();
    var height = $("#canvas").height();

    var cw = 20;
    var d;
    var food;

    var snake_array;

    function paint()
    {
        var ctx = canvas.getContext('2d');
        const number_cell_by_line = 24;
        const size_cell = 20;
        for (var y = 0; y < number_cell_by_line; y++)
            for (var x = 0; x < number_cell_by_line; x++)
                ctx.rect(x * size_cell, y * size_cell, size_cell, size_cell);
        ctx.fillStyle = "rgb(180, 0, 0)";
        ctx.fill();
        ctx.stroke();

        var nx = snake_array[0].x;
        var ny = snake_array[0].y;

        if(d == "right") nx++;
        else if(d == "left") nx--;
        else if(d == "up") ny--;
        else if(d == "down") ny++;

        if(nx == -1 || nx == width/cw || ny == -1 || ny == height/cw || check_collision(nx, ny, snake_array))
        {
            init();
            return;
        }

        if(nx == food.x && ny == food.y)
        {
            var tail = {x: nx, y: ny};
            create_food();
        }
        else
        {
            var tail = snake_array.pop();
            tail.x = nx; tail.y = ny;
        }

        snake_array.unshift(tail);

        for(var i = 0; i < snake_array.length; i++)
        {
            var c = snake_array[i];
            paint_cell(c.x, c.y, "blue");
        }

        paint_cell(food.x, food.y, "yellow");
    }

    function paint_cell(x, y, color)
    {
        ctx.fillStyle = color;
        ctx.fillRect(x*cw, y*cw, cw, cw);
        ctx.strokeStyle = "white";
        ctx.strokeRect(x*cw, y*cw, cw, cw);
    }

    function init()
    {
        d = "right";
        create_snake();
        create_food();

        if(typeof game_loop != "undefined") clearInterval(game_loop);
        game_loop = setInterval(paint, 90);
    }
    init();

    function create_snake()
    {
        var length = 3;
        snake_array = [];
        for(var i = length-1; i>=0; i--)
        {
            snake_array.push({x: i, y:0});
        }
    }

    function create_food()
    {
        food = {
            x: Math.round(Math.random()*(width-cw)/cw),
            y: Math.round(Math.random()*(height-cw)/cw),
        };
    }

    function check_collision(x, y, array)
    {
        for(var i = 0; i < array.length; i++)
        {
            if(array[i].x == x && array[i].y == y)
                return true;
        }
        return false;
    }

    $(document).keydown(function(e){
        var key = e.which;

        if(key == "37" && d != "right") d = "left";
        else if(key == "38" && d != "down") d = "up";
        else if(key == "39" && d != "left") d = "right";
        else if(key == "40" && d != "up") d = "down";
    })
})
