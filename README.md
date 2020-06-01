# Marvelous

The problem statement is as below. Expectation is that the candidate implements this in the Golang.

 

As a game development company, Marvelous limited planned to get new game to market with the existing APIs. There are 3 APIs available

1.       Avengers characters (http://www.mocky.io/v2/5ecfd5dc3200006200e3d64b)

2.       Anti heroes (http://www.mocky.io/v2/5ecfd630320000f1aee3d64d)

3.       Mutants (http://www.mocky.io/v2/5ecfd6473200009dc1e3d64e)

 

Each api gives the list of character and its power level. The power level changes at every 10s interval. Expectation is to call these apis in parallel and store the response in a data structure. Developer has to build a micro-service to do these operations. Also service should expose a api which accepts the name of the character and gives the power level back. Expectation is api should give the latest power level of the character. But the data structure size is limited and can hold upto only 15 characters.

Challenge 1: If there more characters then we should remove the least used characters.

Challenge 2: If there is no space to add incoming characters  until player requests for a character for the first time then drop lowest powered characters to make space.
