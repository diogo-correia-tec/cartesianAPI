# cartesianAPI - Diogo Correia A. Silva 

API server in go that deal with a series of points represented as (x,y) coordinates on a simple 2-dimensional plane. Take a look at https://en.wikipedia.org/wiki/Cartesian_coordinate_system if you need a refresher on this concept.

It has an api route at /api/points that accepts a GET request with the following parameters, and returns a JSON list of points that are within distance from x,y, using the Manhattan distance method. 
The points are beeing returned in order of increasing distance from the search origin.

- x integer (required). This represents the x coordinate of the search origin.
- y integer (required). This represents the y coordinate of the search origin.
- distance integer (required). This represents the Manhattan distance; 

Points within distance from x and y are returned, points outside are filtered out.


## About Manhattan Distance
The Manhattan distance is measured "block-wise", as the distance in blocks between any two points in the plane (e.g. 2 blocks down and 3 blocks over for a total of 5 blocks). 
It is defined as the sum of the horizontal and vertical distances between points on a grid. Formally, where p1 = (x1, y1) and p2 = (x2, y2), distance(p1,p2) = |x1-x2| + |y1-y2|.


### How to Run

Build application
```
go build main.go
```

Run in terminal:
```
go run main.go
```

OR

```
Execute the file "cartesianAPI.exe" or build
```

-You can use your postman or curl to make a request to:

```
GET localhost:8000/api/points?x=100&y=200&distance=100
```

OR 

Enjoy the swagger documentation by openning:
```
http://localhost:8000/api/swagger/index.html
```

*P.S: On startup, the API server reads a list of points from data/points.json.*



