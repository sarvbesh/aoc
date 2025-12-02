#include <bits/stdc++.h>
#include <fstream>
#include <iostream>
#include <string>
using namespace std;

/**
 * Its a circular queue of length 100 starts at 0 and ends at 99
 * Initial pointer is at 50 index
 * 
 * Algorithm:
 * 
 * 1. Read all the move rotations one by one.
 * 2. Calculate the new position after each rotation (moving left or right).
 * 3. Since the dial is circular (99 wraps back to 0), use modulo 100 to handle the wrap-around.
 * 4. After every move, check the new position:
 * 5. If the new position is 0, add 1 to the final count.
 * 6. Return the total count of times the dial hit 0.
 */

class Dial {
public:
  int position = 50;

  void traverseRight(int numberOfTimes, int *count) {
    for (int i = 0; i < numberOfTimes; i++) {
      if (position == 99) {
        position = 0;
      } else {
        position++;
      }
      if (position == 0) {
        *count = *count + 1;
      }
    }
  }

  void traverseLeft(int numberOfTimes, int *count) {
    for (int i = 0; i < numberOfTimes; i++) {
      if (position == 0) {
        position = 99;
      } else {
        position--;
      }
      if (position == 0) {
        *count = *count + 1;
      }
    }
  }
};

int main() {
  ifstream input_file("input.txt");
  string line;
  int count = 0;
  if (!input_file.is_open()) {
    cerr << "Error opening file" << endl;
    return 1;
  }
  char direction;
  int numberOfTimes = 0;

  Dial dial;
  cout << "Initial position: " << dial.position << endl;
  int countOfZero = 0;
  cout << "Initial count: " << countOfZero << endl;
  while (getline(input_file, line)) {
    count++;
    direction = line[0];
    numberOfTimes = stoi(line.substr(1));

    if (direction == 'R') {
      dial.traverseRight(numberOfTimes, &countOfZero);
    } else {
      dial.traverseLeft(numberOfTimes, &countOfZero);
    }
  }

  cout << "Final Position: " << dial.position << endl;
  cout << "Number of lines: " << count << endl;
  cout << "Final count: " << countOfZero << endl;

  input_file.close();
  return 0;
}