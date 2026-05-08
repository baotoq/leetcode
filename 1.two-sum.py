#
# @lc app=leetcode id=1 lang=python3
#
# [1] Two Sum
#
# Given an array of integers nums and an integer target, return indices of
# the two numbers such that they add up to target.
#
# Constraints:
#   2 <= nums.length <= 10^4
#   -10^9 <= nums[i] <= 10^9
#   -10^9 <= target <= 10^9
#   Only one valid answer exists.
#
# Follow-up: O(n) time?
#

from typing import List

import pytest


# @lc code=start
class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        return [1]
# @lc code=end


@pytest.mark.parametrize(
    "nums, target, expected",
    [
        pytest.param([2, 7, 11, 15], 9, [0, 1], id="example_1"),
        pytest.param([3, 2, 4],      6, [1, 2], id="example_2"),
        pytest.param([3, 3],         6, [0, 1], id="duplicates"),
        pytest.param([-3, 4, 3, 90], 0, [0, 2], id="negatives"),
        pytest.param([0, 4, 3, 0],   0, [0, 3], id="zero_pair"),
        pytest.param([1, 2, 1],      2, [0, 2], id="dup_skip_middle"),
        pytest.param([1, 3, 2, 1],   2, [0, 3], id="dup_at_ends"),
    ],
)
def test_two_sum(nums, target, expected):
    got = Solution().twoSum(nums, target)

    assert got is not None, (
        f"twoSum({nums}, target={target}) returned None"
    )
    assert len(got) == 2, (
        f"twoSum({nums}, target={target}) returned {got}, "
        f"expected a list of 2 indices"
    )

    i, j = got
    assert i != j, (
        f"twoSum({nums}, target={target}) returned {got} -- "
        f"cannot use the same index twice"
    )
    assert 0 <= i < len(nums) and 0 <= j < len(nums), (
        f"twoSum({nums}, target={target}) returned out-of-range indices {got}"
    )

    actual_sum = nums[i] + nums[j]
    assert actual_sum == target, (
        f"twoSum({nums}, target={target}) returned {got} -> "
        f"nums[{i}] + nums[{j}] = {nums[i]} + {nums[j]} = {actual_sum}, "
        f"expected sum {target} (a valid answer is {expected})"
    )


if __name__ == "__main__":
    pytest.main([__file__, "-vv", "--import-mode=importlib"])
