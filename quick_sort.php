<?php

function quickSort(array $arr) {
	if (count($arr) < 2) 
		return $arr;

	$left = 0;
	$right = count($arr) - 1;

	foreach ($arr as $key => $value) {
		if ($value < $arr[$right]) {
			$arr[$key] = $arr[$left];
			$arr[$left] = $value;
			$left++;
		}
	}

	$tmp = $arr[$right];
	$arr[$right] = $arr[$left];
	$arr[$left] = $tmp;

	// $l = quickSort(array_slice($arr, 0, $left));
	// $r = quickSort(array_slice($arr, $left));
	// return array_merge($l, $r);	

	// quick
	$l = quickSort(array_slice($arr, 0, $left));
	$r = quickSort(array_slice($arr, $left+1));
	return array_merge($l, $value, $r);	
}

$res = quickSort([3,7,2,4,5]);
print_r($res);


function quickSort2($arr) {
	if(count($arr) < 2) 
		return $arr;

	$base = $arr[0];
	$left = $right = [];

	foreach ($arr as $key => $value) {
		if ($value < $base) {
			$left[] = $value;
		}else {
			$right[] = $value;
		}
	}

	array_shift($right);


	$left = foo($left);
	$right = foo($right);

	return array_merge($left, [$base], $right);
}
