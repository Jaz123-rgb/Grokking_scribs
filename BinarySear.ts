function binaryearch(nums: number[], target: number): number{
    let left: number = 0;
    let rigth: number = nums.length - 1

    while(left <= rigth){
        
	const mid: number = Math.floor((left + rigth)/2);

	if(nums[mid] === target) return mid;

	if(target< nums[mid]) rigth = mid - 1;
        else left = mid + 1;
    }
    return - 1
}

let my_list:number[];

my_list = [1,2,3,4,5,6];

let rannum: number = 6;

console.log(binaryearch(my_list, rannum));