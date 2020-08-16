<?php

function Foo() {
  // int
  $a = 100;
  // float
  $b = 1.5;
  // string
  $c = "Hello";
  // bool
  $d = true;
  // output
  echo $a;
  echo $b;
  echo $c;
  echo $d;
  // if-else
  if ($a == 100) {
    // variable inside
    $e = 10;
  } else {
    // and here, so variable should be inside
    $e = 10;
  }
  // output this variable
  echo $e;
  // another if-else
  if ($a == 100) {
    // variable with int type
    $f = 10;
  } else {
    // here is string
    $f = "10";
  }
  // so $f has type int|string
  echo $f;
  // support only single-type array
  $f = [1,2,3];
  echo $f;
  // fetch by index
  echo $f[1];
  // assign by index
  $f[1] = 10;
  echo $f;
  // simple array
  $g = [1,2,3];
  echo $g;
  // adding element
  $g[] = 100;
  echo $g;
  // and associative array with single-type keys
  $f = ["Key1" => 1, "Key2" => 2, "Key3" => 3];
  echo $f;
  // fetch by key
  echo $f["Key1"];
  // assign by key
  $f["Key1"] = 5;
  echo $f;
  // while
  $i = 0;
  while ($i < 20) {
    echo $i;
    $i++;
  }
  // for
  for ($i = 0; $i < 20; $i++) {
    echo $i + 5;
  }
  $qw = 1.5;
  // different operators
  echo $qw + 5 - 56.56 * 6 / 56;
}