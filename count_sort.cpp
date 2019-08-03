/*
 * 在保证全是非负的情况使用，并且最大的数字不是特别大
 * 
 * */

#include <iostream>
#include <algorithm>

void count_sort(int* array, int len) {
    int* output = new int[len+1];
    int max_n = *std::max_element(array, array+len);
    int* store = new int[max_n+1]();
    for(int i=0; i<len; i++) {
        store[array[i]]++;
    }
    for(int i=1; i<=max_n; i++) {
        store[i] += store[i-1];
    }
    for(int i=len-1; i>=0; i--) {
        output[store[array[i]]] = array[i];
        store[array[i]] --;
    }
    for(int walker=1; walker<=len; walker++) {
        std::cout << output[walker] << " ";
    }
    std::cout << std::endl;
    delete []store, delete []output;
}

int main()
{
    int arrar[] = {2,5,3,0,2,3,0,3};
    count_sort(arrar, 8);
    int arr[] = {4,1,9,10};
    count_sort(arr, 4);
    return 0;
}

