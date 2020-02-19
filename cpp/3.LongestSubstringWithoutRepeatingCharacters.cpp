///
/// LeetCode 
/// By Reza Ebrahimi <reza.ebrahimi.dev@gmail.com>
///
/// 3. Longest Substring Without Repeating Characters
/// https://leetcode.com/problems/longest-substring-without-repeating-characters
///
/// Time Submitted | Status   | Runtime | Memory | Language
/// 2019/04/18     | Accepted | 16 ms	| 9.4 MB | cpp
///
/// Runtime: 16 ms, faster than 92.75% of C++ online submissions for Longest Substring Without Repeating Characters.
/// Memory Usage: 9.4 MB, less than 99.46% of C++ online submissions for Longest Substring Without Repeating Characters.
///


class Solution {
public:
    int lengthOfLongestSubstring(string s)
    {
        int len = 0;
        vector<char> sub = {};

        for (auto& c : s) {
            auto it = find(sub.begin(), sub.end(), c);
            if (it != sub.end()) {
                sub.erase(sub.begin(), next(it));
            }
            sub.push_back(c);
            len = max(len, (int)sub.size());
        }

        return len;
    }
};
