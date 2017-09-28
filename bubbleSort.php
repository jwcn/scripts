<?php

// 冒泡排序

function bubbleSort($arr) {
	$isOver = false;
	$n = count($arr) - 1;

	while (!$isOver) {
		$exchange = false;
		for ($i=0; $i < $n; $i++) {
			if ($arr[$i] > $arr[$i+1]) {
				$tmp = $arr[$i+1];
				$arr[$i+1] = $arr[$i];
				$arr[$i] = $tmp;


				$exchange = true;	// 有交换过
			}
		}

		if (!$exchange) {	// 一次没换过, 结束循环和排序
			$isOver = true;
		}

		$n--;
	}

	return $arr;
}

$res = bubbleSort([3,7,2,4,5]);
print_r($res);

/*
1. 从右侧第一位开始与左侧每一位对比,小就换
2. 第二次循环对比排除最左侧第一位, 第三次排除第二位, 依次类推
*/
function foo(array $arr) {
	if (count($arr) < 2)
		return $arr;

	$len = count($arr);
	for ($i = 0; $i < $len; $i++) {
		for ($j = $len-1; $j > $i; $j--) {
			if ($arr[$j] < $arr[$j-1]) {
				$tmp = $arr[$j];
				$arr[$j] = $arr[$j-1];
				$arr[$j-1] = $tmp;
			}
		}
	}

	return $arr;
}
