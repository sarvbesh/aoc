#include <bits/stdc++.h>
#include <fstream>
#include <iostream>
#include <string>
using namespace std;

/**
 * The dial is circular with positions 0-99
 * Initial position is at 50
 *
 * Goal: Count how many times the dial points at 0 after completing each rotation
 *
 * - R means move right (toward higher numbers, wraps 99->0)
 * - L means move left (toward lower numbers, wraps 0->99)
 *
 * Algorithm:
 * 1. Read each line from input file
 * 2. Parse direction (L/R) and distance
 * 3. Calculate new position using modular arithmetic
 * 4. Check if final position == 0, if so increment count
 * 5. Return total count
 */

class Dial {
public:
  int position = 50;

  void traverseRight(int numberOfTimes, int *count) {
    position = (position + numberOfTimes) % 100;
    if (position == 0) {
      *count = *count + 1;
    }
  }

  void traverseLeft(int numberOfTimes, int *count) {
    position = (position - numberOfTimes % 100 + 100) % 100;
    if (position == 0) {
      *count = *count + 1;
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
    // cout << "Line: " << count << " Direction: " << direction
    //    << " Change in position: " << numberOfTimes << endl;

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