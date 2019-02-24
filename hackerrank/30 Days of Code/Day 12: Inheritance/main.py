class Person:
	def __init__(self, firstName, lastName, idNumber):
		self.firstName = firstName
		self.lastName = lastName
		self.idNumber = idNumber
	def printPerson(self):
		print "Name:", self.lastName + ",", self.firstName
		print "ID:", self.idNumber

class Student(Person):
    def __init__(self, firstName = None, lastName = None, idNum = 0, scores = []):
        self.firstName = firstName
        self.lastName = lastName
        self.idNumber = idNum
        self.scores = scores

    def calculate(self):
        self.scoresCount = 0
        for i in self.scores:
            self.scoresCount += i

        self.avgScore = self.scoresCount / len(self.scores)
        self.grade = None
        if 90 <= self.avgScore <= 100:
            self.grade = "O"
        elif 80 <= self.avgScore < 90:
            self.grade = "E"
        elif 70 <= self.avgScore < 80:
            self.grade = "A"
        elif 55 <= self.avgScore < 70:
            self.grade = "P"
        elif 40 <= self.avgScore < 55:
            self.grade = "D"
        elif self.avgScore < 40:
            self.grade = "T"
        
        return self.grade

line = raw_input().split()
firstName = line[0]
lastName = line[1]
idNum = line[2]
numScores = int(raw_input()) # not needed for Python
scores = map(int, raw_input().split())
s = Student(firstName, lastName, idNum, scores)
s.printPerson()
print "Grade:", s.calculate()
