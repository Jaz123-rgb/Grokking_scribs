function binaryearch(nums, target) {
    var left = 0;
    var rigth = nums.length - 1;
    while (left <= rigth) {
        var mid = Math.floor((left + rigth) / 2);
        if (nums[mid] === target)
            return mid;
        if (target < nums[mid])
            rigth = mid - 1;
        else
            left = mid + 1;
    }
    return -1;
}
var my_list;
my_list = [1, 2, 3, 4, 5, 6];
var rannum = 6;
console.log(binaryearch(my_list, rannum));
