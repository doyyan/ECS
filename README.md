## Author: Ravikumar Sivaraj 
###### Dated: February 2021
&nbsp;


The goal of this is to create a Reusable Basket-Pricer component written in Go.

**The Pricer** implements an interface that takes in an type named ```Basket```.  It calculates the discounts, subtotals and totals and prints out the reuslts AND returns a Reciept object that has the following fields:-

   A itemised ```Receipt``` with each line having an

   ```Item Name  | Original Price (without Discount) | Applicable Discount | Discounted Price (if applicable)```

   Prices are returned as floating point numbers with width 8, precision 2
&nbsp;

## To run a Sample
Run the ./run.sh in the parent directory in a unix shell. It has calculations for 2 products, Biscuits and Baked Beans with the former priced at £1.20 each with no discounts and the latter at £0.99 with a Buy 3 get 1 free. Open the shoppingtrip.go file and change the Number of items and test.
&nbsp;
&nbsp;

### Types and Structure

The ```Basket``` implements a ```Pricer``` which calculates the price of the items in the Shopping Basket.
Each ```Basket``` has a ```Catalogue``` attached to it which contains a list of ```Product``` with Pricing and ```Offer``` information.

### Remote `Offer` Teams 

The ```Catalogue``` exposes options for teams to Update ```Offers``` for ```Products```




