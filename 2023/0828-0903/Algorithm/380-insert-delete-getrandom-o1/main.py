import random

class RandomizedSet:
    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.data = []
        self.index = {}
        self.length = 0


    def insert(self, val: int) -> bool:
        """
        Inserts a value to the set. Returns true if the set did not already
        contain the specified element.
        """
        if val in self.index:
            return False
        self.data.append(val)
        self.index[val] = self.length
        self.length += 1
        return True


    def remove(self, val: int) -> bool:
        """
        Removes a value from the set. Returns true if the set contained the
        specified element.
        """
        if val not in self.index:
            return False
        index = self.index[val]
        self.data[index] = self.data[-1]
        self.index[self.data[-1]] = index
        self.data.pop()
        self.length -= 1
        del self.index[val]
        return True


    def getRandom(self) -> int:
        """
        Get a random element from the set.
        """
        return self.data[random.randint(0, self.length - 1)]



# Your RandomizedSet object will be instantiated and called as such:
# obj = RandomizedSet()
# param_1 = obj.insert(val)
# param_2 = obj.remove(val)
# param_3 = obj.getRandom()

if __name__ == '__main__':
    obj = RandomizedSet()
    print(obj.insert(1))
    print(obj.remove(2))
    print(obj.insert(2))
    print(obj.getRandom())
    print(obj.remove(1))
    print(obj.insert(2))
    print(obj.getRandom())