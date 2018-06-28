<?php

Class Node 
{
	public $data = null;
	protected $leftNode = null;
	protected $rightNode = null;
	protected $parentNode = null;

	public function __construct ($value) {
		$this->data = $value;
	}

	public function setLeft($node) {
		$this->leftNode = $node;
		$node->setParent($this);
	}

	public function getLeft() {
		return $this->leftNode;
	}

	public function setRight($node) {
		$this->rightNode = $node;
		$node->setParent($this);
	}

	public function getRight() {
		return $this->rightNode;
	}

	public function setParent($node) {
		$this->parentNode = $node;
	}

	public function getParent() {
		return $this->parentNode;
	}
}

Class Tree 
{
	public static function insert($root, $node) {
		if ($node->data < $root->data) {
			if ($root->getLeft()) {
				self::insert($root->getLeft(), $node);
			} else {
				$root->setLeft($node);
			}
		} else {
			if ($root->getRight()) {
				self::insert($root->getRight(), $node);
			} else {
				$root->setRight($node);
			}
		}
	}

	public static function find($root, $value) {
		if ($value < $root->data) {
			return self::find($root->getLeft(), $value);
		} else if ($value > $root->data) {
			return self::find($root->getRight(), $value);
		} else {
			return $root;
		}
	}

	public static function findMin($root) {
		if ($root->getLeft()) {
			return self::findMin($root->getLeft());
		}

		return $root;
	}

	public static function findMax($root) {
		if ($root->getRight()) {
			return self::findMax($root->getRight());
		}

		return $root;
	}
}

$root = new Node(20);
Tree::insert($root, new Node(15));
Tree::insert($root, new Node(22));
Tree::insert($root, new Node(13));
Tree::insert($root, new Node(18));

var_dump($root);
// var_dump(Tree::find($root, 13));exit;
// var_dump(Tree::findMin($root));
// var_dump(Tree::findMax($root));
