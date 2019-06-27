<?php

use \Psr\Http\Message\ServerRequestInterface as Request;
use \Psr\Http\Message\ResponseInterface as Response;

$app = new \Slim\App;
// Run app

$client = new Restaurant\RestaurantServiceClient('localhost:50051', [
    'credentials' => Grpc\ChannelCredentials::createInsecure(),
]);

// Utilities
function mapResponseToPHPObject($r) {
    $obj = [
        "rest_id" => $r->getRestId(),
        "name" => $r->getName(),
        "rating" => $r->getRating(),
        "cuisines" => $r->getCuisines(),
        "address" => $r->getAddress(),
        "timing" => $r->getTiming(),
        "cft" => $r->getCft()
    ];

    return $obj;
}

function mapRequestToRestaurantObject($r) {
    $restaurant = new Restaurant\Restaurant;
    
    $restaurant->setName($r['name']);
    $restaurant->setRating($r['rating']);
    $restaurant->setCuisines($r['cuisines']);
    $restaurant->setAddress($r['address']);
    $restaurant->setTiming($r['timing']);
    $restaurant->setCft($r['cft']);

    return $restaurant;
}


// Define app routes
$app->group('/api', function () use ($app){
    $app->get('/restaurants', function (Request $request, Response $response, array $args) {
        global $client;
    
        $empty = new Restaurant\PBEmpty;
        $call = $client->GetAllRestaurants($empty);
        $restaurants = $call->responses();
        
        $final = [];
        foreach($restaurants as $rest) {
            $final[] = mapResponseToPHPObject($rest);
        }
        return $response->withJson($final);
    });
    
    $app->get('/restaurant/{rest_id}', function (Request $request, Response $response, array $args) {
        global $client;
    
        $rest_id = new Restaurant\RestaurantId;
        $rest_id->setRestId($args['rest_id']);
    
        list($res, $status) = $client->ReadRestaurant($rest_id)->wait();
        
        $final = mapResponseToPHPObject($res);
        return $response->withJson($final);
    });
    
    $app->post('/restaurant', function(Request $request, Response $response, array $args){
        global $client;
    
        $req = $request->getParsedBody();
        $restaurant = mapRequestToRestaurantObject($req);
    
        list($res, $status) = $client->CreateRestaurant($restaurant)->wait();
    
        $final = ['rest_id' => $res->getRestId()];
    
        return $response->withJson($final);
    });
    
    $app->put('/restaurant', function (Request $request, Response $response, array $args) {
        global $client;
    
        $r = $request->getParsedBody();
        $restaurant = new Restaurant\Restaurant;

        $restaurant->setRestId($r['rest_id']);
        $restaurant->setName($r['name']);
        $restaurant->setRating($r['rating']);
        $restaurant->setCuisines($r['cuisines']);
        $restaurant->setAddress($r['address']);
        $restaurant->setTiming($r['timing']);
        $restaurant->setCft($r['cft']);
    
        list($res, $status) = $client->UpdateRestaurant($restaurant)->wait();
    
        $final = ['status' => $res->getSuccess()];
        return $response->withJson($final);
    });
    
    $app->delete('/restaurant/{rest_id}', function (Request $request, Response $response, array $args) {
        global $client;
    
        $rest_id = new Restaurant\RestaurantId;
        $rest_id->setRestId($args['rest_id']);
    
        list($res, $status) = $client->DeleteRestaurant($rest_id)->wait();
        
        $final = ['status' => $res->getSuccess()];
        return $response->withJson($final);
    });
});