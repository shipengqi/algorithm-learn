---
title: 复杂度分析
weight: 1
---

## 什么是算法复杂度？

**算法复杂度**（Algorithm Complexity）是指算法在运行过程中消耗的资源（通常是**时间**或**空间**）的量级。简单来说，它告诉我们一个算法在处理不同规模的数据时，会变得多“慢”或多“快”。

假设有两个算法，都用于排序一个数组。对于一个包含 10 个数字的数组，两个算法可能都能在瞬间完成排序。但如果数组中有 100 万个数字，其中一个算法可能需要几分钟，而另一个只需要几秒钟。这就是算法复杂度的体现。

## 为什么需要复杂度分析

在实际工作中，我们把代码跑一遍，通过统计、监控，就能得到算法执行的时间和占用的内存大小。为什么还要做时间、空间复杂度分析？

因为这种方法叫**事后统计法**，这种方法的局限性：

1. 依赖测试环境。代码在不同的环境运行，结果是不同的，比如一个酷睿 i9，和酷睿 i3，很明显 i9 处理速度要快的多。
2. 测试结构受数据规模的影响。测试数据规模太小，测试结果可能无法真实地反应算法的性能。

## 如何衡量算法复杂度

算法复杂度通常用大 `O` 记号来表示。大 O 记号描述了算法在最坏情况下所需的**时间或空间**与**输入规模**的关系。以下是一些常见的复杂度类型：

### O(1)：常数时间复杂度

无论输入多大，操作次数都是固定的。

```javascript
function add(a, b) {
    return a + b; // 单次操作
}
```

无论 a 和 b 的值是多少，操作次数始终为 1，因此这个函数的时间复杂度是 `O(1)`。

### O(n)：线性时间复杂度

如果算法的运行时间与输入规模成正比，那么它的复杂度是 `O(n)`。常见于单层循环。例如，遍历一个数组并打印每个元素：

```javascript
function printArray(arr) {
    for (let i = 0; i < arr.length; i++) {
        console.log(arr[i]);
    }
}
```

如果数组长度为 10，需要执行 10 次操作；长度为 100，则需要 100 次操作。因此，这个函数的时间复杂度是 `O(n)`。

### O(n^2)：平方时间复杂度

如果算法的运行时间与输入规模的平方成正比，复杂度为 `O(n^2)`。常见于嵌套循环：

```javascript
function printPairs(arr) {
    for (let i = 0; i < arr.length; i++) {
        for (let j = 0; j < arr.length; j++) {
            console.log(arr[i], arr[j]);
        }
    }
}
```

如果数组长度为 10，需要执行 `10×10 = 100` 次操作；长度为 100，则需要 `100×100 = 10,000` 次操作。这种复杂度的算法在处理大数据时会变得非常慢。

### O(logn)：对数时间复杂度

对数复杂度，常见于“分而治之”的算法，如二分查找等场景。例如，查找一个有序数组中的某个元素：

```javascript
function binarySearch(arr, target) {
    let left = 0;
    let right = arr.length - 1;
    while (left <= right) {
        const mid = Math.floor((left + right) / 2);
        if (arr[mid] === target) {
            return mid;
        } else if (arr[mid] < target) {
            left = mid + 1;
        } else {
            right = mid - 1;
        }
    }
    return -1;
}
```

每次循环都将搜索范围减半，因此时间复杂度为 `O(logn)`。

大 O 时间复杂度实际上并不具体表示代码真正的执行时间，而是**表示代码执行时间随数据规模增长的变化趋势**，所以，也叫作**渐进时间复杂度（asymptotic time complexity），简称时间复杂度**。

### O(nlogn)：线性对数时间复杂度

如果算法的运行时间与输入规模的乘积成正比，复杂度为 `O(nlogn)`。常见于排序算法，如归并排序和快速排序。

```javascript
function mergeSort(arr) {
    if (arr.length <= 1) {
        return arr; // 如果数组只有一个元素，直接返回
    }

    // 分解：将数组分成两半
    const mid = Math.floor(arr.length / 2);
    const left = mergeSort(arr.slice(0, mid)); // 对左半部分排序
    const right = mergeSort(arr.slice(mid));  // 对右半部分排序

    // 合并：将两个有序数组合并成一个有序数组
    return merge(left, right);
}

function merge(left, right) {
    const result = [];
    let i = 0; // 左数组指针
    let j = 0; // 右数组指针

    // 比较两个数组的元素，将较小的元素放入结果数组
    while (i < left.length && j < right.length) {
        if (left[i] <= right[j]) {
            result.push(left[i]);
            i++;
        } else {
            result.push(right[j]);
            j++;
        }
    }

    // 将剩余的元素直接添加到结果数组
    return result.concat(left.slice(i)).concat(right.slice(j));
}

const arr = [38, 27, 43, 3, 9, 82, 10];
const sortedArr = mergeSort(arr);
console.log(sortedArr); // 输出：[3, 9, 10, 27, 38, 43, 82]
```

每次递归将数组分成两半（对数级别），合并操作需要线性时间。因此，总时间复杂度为 `O(n log n)`。

### 如何分析一段代码的时间复杂度

1. 关注循环、递归和函数调用等操作。记录每种操作的执行次数。
2. 只关注最高阶项，通常忽略低阶项和常数。例如，`O(n² + n)` 简化为 `O(n²)`，`O(3n)` 简化为 `O(n)`。

示例：

```go
func cal3(n int) int {
  ret := 0
  for i := 1; i <= n; i++ {
    ret = ret + f(i)
  }
  return ret
}

func f(n int) int {
  sum := 0
  for i := 1; i <= n; i++ {
    sum = sum + i
  }
  return sum
}
```

单独看 `cal3()` 函数。假设 `f()` 只是一个普通的操作，循环执行 n 次，那么时间复杂度是 `O(n)`。但 `f()` 函数里面也是一个循环，它的时间复杂度也是 `O(n)`，所以，整个 `cal()` 函数的时间复杂度就是  `O(n*n) = O(n^2)`。

示例 2：

```javascript
function example(n) {
    for (let i = 0; i < n; i++) {
        console.log(i); // O(n)
    }
    for (let j = 0; j < n * n; j++) {
        console.log(j); // O(n²)
    }
}
```

- 第一个循环的复杂度是 `O(n)`。
- 第二个循环的复杂度是 `O(n^2)`。
- 总复杂度是 `O(n + n^2)`。
- 忽略低阶项 `O(n)` 后最终复杂度是：`O(n^2)`。

### 复杂度量级

时间复杂度可以分为**多项式量级**和**非多项式量级**。

常见的多项式量级：

- `O(1)`：常数时间复杂度
- `O(n)`：线性时间复杂度
- `O(n log n)`：线性对数时间复杂度
- `O(n^2)`：二次时间复杂度
- `O(n^3)`：三次时间复杂度
- `O(n^k)`：k 次多项式时间复杂度

多项式量级的算法通常被认为是“可解”的，因为它们的运行时间随着输入规模的增加而以多项式的方式增长，这在实际应用中通常是可接受的。

常见的非多项式量级

- `O(2^n)`：指数时间复杂度
- `O(n!)`：阶乘时间复杂度

非多项式量级的算法通常被认为是“不可解”的，因为它们的运行时间随着输入规模的增加而以指数级或阶乘级的方式增长，这在实际应用中通常是不可接受的。

指数时间复杂度示例：

```javascript
function fibonacci(n) {
    if (n <= 1) {
        return n;
    }
    return fibonacci(n - 1) + fibonacci(n - 2);
}
```

每次递归调用会生成两个新的调用，调用次数呈指数增长，因此时间复杂度为 `O(2ⁿ)`。

阶乘时间复杂度示例：

```javascript
function permutations(arr) {
    if (arr.length === 1) {
        return [arr];
    }
    const result = [];
    for (let i = 0; i < arr.length; i++) {
        const current = arr[i];
        const remaining = arr.slice(0, i).concat(arr.slice(i + 1));
        const remainingPermutations = permutations(remaining);
        for (let perm of remainingPermutations) {
            result.push([current].concat(perm));
        }
    }
    return result;
}
```

每次递归调用都会生成 n 个新的调用，调用次数呈阶乘增长，因此时间复杂度为 `O(n!)`。
