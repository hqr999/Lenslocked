## Exercises

We haven't used MVC in our code, so it would be hard to do any coding exercises. Instead, we are going to just review a few questions and answers related to MVC.


### Ex1 - What does MVC stand for?

What does each letter in the MVC acronym stand for?

MVC stands for MODEL, VIEW AND CONTROLLER.

### Ex2 - What is each layer of MVC responsible for?

Try to jot down what each layer of MVC is responsible for and review the chapter to see if you are correct.

MODEL-Is mainly responsible for the "hard" work. It deals with the database and other data that comes out of different APIs.


VIEW-It is responsible for rendering the data. This data could be HTML or JSON that comes from a REST API. It shouldnt have much logic as well.


CONTROL-The name is self explanatory, its responsability is mainly with regards with controlling how the application is gonna work, basically being responsible for when methods from model and view are gonna be called.


### Ex3 - What are some benefits and disadvantages to using MVC?

Take a moment to think about why MVC might help keep your code organized better than an ad-hoc structure.

Why might it be easier for others to help you work on your code?

Why is it easier to use when you aren't 100% certain what your final application will look like?

MVC is a good choide because a lot of frameworks make use it 
and thus greatly facilitating the start of an application.Frameworks like:Ruby on Rails(Ruby),Django(Python) and Laravel(PHP) are good examples of it.

Can you think of any reasons why MVC might not be a great fit for some projects?

For projects with large amounts of different libraries,MVC simplistic structure simply isnt enough to abstract the depth of this complex application.

### Ex4 - Read about other ways to structure code

Another exercise worth doing is to read up on other code structures.

Ben Johnsons <gobeyond.dev> has some interesting articles worth reading as well, but some of them may use code you don't quite understand at this point, so don't feel like you need to grasp it all to proceed.

Kat Zien also has a talk on this from a past Gophercon at <https://www.youtube.com/watch?v=oL6JBUk6tj0>

