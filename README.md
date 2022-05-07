Exercise for Friday May 6th.
I recommend people pair up for this exercise, or even do it as a full group exercise. I estimate this will take about 45 minutes to an hour to complete.

- [x] Create a custom type (drugid) from an int to represent a drug.
- [x] Add a type method to drugid that prints out the name of the drug, using a global prepopulated map where drugid is the key, and the name of the drug is the value
- [x] Create two pharmacies that implement the following methods
  - DispenseDrug(drug drugid) error
  - LoadInventory(drug drugid, quantity int)
  - CheckInventory(drug drugid) (int, error)
  - DisplayFullInventory()
  - These methods should printout the action taken, and the id and name of the drug being operated on
  - Each pharmacy should be able to store the inventory of drugs as a map with the drugID as the key and the quantity as the value. DispenseDrug and LoadInventory should update the inventory accordingly.
  - One pharmacy has the method to process insurance, the other is cash only.
- [ ] Create a pharmacy router, that contains references to both pharmacies, and has a method called
ProcessOrder(drug drugid) error
  - This method will choose the pharmacy that as the specified drug in its inventory, if both pharmacies have the drug in inventory it shall chose the one with the largest inventory. If they are the same, it shall pick one at random. If it is unable to process the drug, it shall return an error.
- [ ] In main, load up the pharmacies and then exercise the Processing of the drugs.
- [ ] Make sure there is full test coverage and Benchmarking.
