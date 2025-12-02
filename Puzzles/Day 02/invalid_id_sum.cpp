#include <bits/stdc++.h>
#include <fstream>
#include <iostream>
#include <sstream>
#include <stdexcept>
#include <string>
#include <sys/types.h>
#include <unordered_map>
using namespace std;

/**
 * Input: ID strings format: firstID-LastID at each line (ranges)
 * Goal: Identify Invalid ID's and sum up all those
 * Invalid IDs has numbers with a frequency of atleast 2
 * For example: 55 (5 twice), 6464 (64 twice), and 123123 (123 twice)
 *
 *
 * My Approach:
 * 1. Read the entire line
 * 2. Create an Integer (int) vector of pairs where each pair stores [{firstID,lastID},{},{}......]
 * 3. First create for the entire data. Lastly, for each pair
 * 4. Check the pattern:
 * 5. If it's of even length and the first substring and second sunstring match then that's INVALID
 * 6. Return that value & add it to the sum. Else, return 0
 */

vector<string> split(const string &s, char delimiter) {
  vector<string> tokens;
  string token;
  istringstream tokenStream(s);

  while (getline(tokenStream, token, delimiter)) {
    tokens.push_back(token);
  }

  return tokens;
}

pair<string, string> getIDPair(string s) {
  pair<string, string> ids;
  string token;
  stringstream ss(s);
  int count = 0;
  while (getline(ss, token, '-')) {
    if (count == 0) {
      try {
        ids.first = token;
      } catch (const invalid_argument &e) {
        throw runtime_error("Invalid id: not a number");
      }
    } else if (count == 1) {
      try {
        ids.second = token;
      } catch (const invalid_argument &e) {
        throw runtime_error("Invalid id: not a number");
      }
    }
    count++;
  }
  return ids;
}

/**
 * Part 2: ID is invalid if it is made only of some sequence of digits repeated at least twice
 * For example: 12341234 (1234 two times), 123123123 (123 three times), 
 * 1212121212 (12 five times), and 1111111 (1 seven times)
 * */ 

long long getInvalidID(long long id) {
  string s = to_string(id);
  int length = s.length();

  for (int patternLength = 1; patternLength <= length / 2; patternLength++) {
    if (length % patternLength != 0)
      continue;

    string pattern = s.substr(0, patternLength);

    string repeated = "";
    int repeatedTimes = length / patternLength;

    for (int i = 0; i < repeatedTimes; i++) {
      repeated += pattern;
    }

    if (s == repeated) {
      return id;
    }
  }
  return 0;
}

int main() {
  long long sum = 0;

  ifstream input_file("input.txt");
  string line;

  if (!input_file.is_open()) {
    cerr << "Error opening file" << endl;
    return 1;
  }

  getline(input_file, line);
  char delimiter = ',';

  vector<string> inputVector = split(line, delimiter);
  // for (string s : inputVector) {
  //   cout << s << endl;
  // }

  for (string s : inputVector) {
    pair<string, string> idRange = getIDPair(s);
    // cout << idRange.first << " " << idRange.second << endl;
    long long tempFirst = stoll(idRange.first);
    long long tempSecond = stoll(idRange.second);
    // cout << tempFirst << " " << tempSecond << endl;

    for (long long i = tempFirst; i <= tempSecond; i++) {
      sum += getInvalidID(i);
      // cout << getInvalidID(i) << endl;
    }
  }

  cout << "Invalid ID sum: " << sum << endl;
  return 0;
}

/**
 * Output: 
 * Part 1: 21139440284
 * Part 2: 38731915928
 */