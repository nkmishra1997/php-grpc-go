<?php

// from task: protoc --proto_path=proto  --php_out=php_client --grpc_out=php_client --plugin=protoc-gen-grpc=/Users/zomato/grpc/bins/opt/grpc_php_plugin proto/restaurant.proto

require dirname(__FILE__).'/vendor/autoload.php';

require dirname(__FILE__). '/routes/routes.php';


$app->run();



// Non Http PHP Client
/*
$restaurant = new Restaurant\Restaurant;
$restaurant->setName('BurgerKing');
$restaurant->setRating(4.8);
$restaurant->setCuisines('Fast Food, Burgers');
$restaurant->setAddress('CyberHub, Gurgaon');
$restaurant->setTiming('11:00 AM - 11:00 PM');
$restaurant->setCft(250);

list($res1, $status) = $client->CreateRestaurant($restaurant)->wait();
echo sprintf("Restaurant Created %s\n", $res1->getRestId());

list($res2, $status) = $client->ReadRestaurant($res1)->wait();
var_dump($res2);
echo sprintf("Restaurant Read %s\n", $res2->getName());

$restaurant->setRating(4.32);
$restaurant->setCuisines('Bad Unhealthy Food, Cold Coffee');

list($res3, $status) = $client->UpdateRestaurant($restaurant)->wait();
var_dump($res3);

list($res4, $status) = $client->DeleteRestaurant($res1)->wait();
var_dump($res4);
*/