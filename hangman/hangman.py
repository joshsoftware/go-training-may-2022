# hangman.py
#hangman game implementation in python

import random

#create word list
wordList = [
    ("keyboard", "Used to Type")  , 
    ("mouse","Used to scroll"),
    ("laptop","Used to code"),
    ("table","a piece of furniture with a flat top used for writing"),
    ("chair","Something you sit on"),
    ("headphones","Used to listen to music"),
]

chances = 1

 
exit = "y"
while exit != "n":
    while chances > 0 :
        #Randomly pick one word  
        randomChoice = random.randrange(len(wordList))
        word = wordList[randomChoice][0]

        #Set chances to len+1
        chances = len(word) +1
        blankList = ["_"] *len(word)
        print()
        print()
        print("HINT: ", wordList[randomChoice][1])
        for i in range(chances):
            for i in blankList:
                print(i, end=" ")
            print()
            print()
            print("There are ", chances, "chances left.")
            print()
            guess = input("Please enter a letter to guess! \n")
            print()
            numOccurence = word.count(guess)
            wordSliced = word
            lengthSlicedOut = 0
            for i in range(numOccurence):
                blankList[wordSliced.find(guess) + lengthSlicedOut] = guess
                lengthSlicedOut += len(word[:numOccurence+1])
                if numOccurence >1:
                    wordSliced = word[numOccurence+1:]
            if blankList == list(word):
                print("Correct! The word was ",word )
                break
            
            chances = chances-1  

    if blankList != list(word):
        print("Oops! The word was ",word )
        print("Try again!")
    exit = input("Do you wish to play again?(y/n)")
    chances = 1

                

