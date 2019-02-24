class Difference:
    def __init__(self, a):
        self.__elements = a

    def computeDifference(self):
        self.maxDif = []
        for i, v in enumerate(self.__elements):
            for j in self.__elements[i+1:]:
                self.diff = v - j
                if self.diff < 0:
                    self.diff = -self.diff

                self.maxDif.append(self.diff)

        self.maximumDifference = sorted(self.maxDif, reverse=True)[0]


_ = raw_input()
a = [int(e) for e in raw_input().split(' ')]

d = Difference(a)
d.computeDifference()

print d.maximumDifference
