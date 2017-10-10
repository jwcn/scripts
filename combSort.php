<?php

echo '<pre>';
$res = foo([3,7,2,4,5]);
print_r($res);

function combSort($items) {
	$swapped   = true;
	$shrink = 1.3;
	$counter = $gap = count($items);

	while ($swapped) {
		$swapped = false;

		$gap = intval($gap / $shrink);
		for ($i = 0; $i+$gap < $counter; $i++) {
			if ($items[$i] > $items[$i+$gap]) {
				$temp = $items[$i];
				$items[$i] = $items[$i+$gap];
				$items[$i+$gap] = $temp;
				$swapped = true;
			}
		}
	}

	return $items;
}