<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Restaurant;

/**
 */
class RestaurantServiceClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * Create a new restaurant.
     * @param \Restaurant\Restaurant $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function CreateRestaurant(\Restaurant\Restaurant $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/restaurant.RestaurantService/CreateRestaurant',
        $argument,
        ['\Restaurant\RestaurantId', 'decode'],
        $metadata, $options);
    }

    /**
     * Read a given restaurant by rest_id
     * @param \Restaurant\RestaurantId $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function ReadRestaurant(\Restaurant\RestaurantId $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/restaurant.RestaurantService/ReadRestaurant',
        $argument,
        ['\Restaurant\Restaurant', 'decode'],
        $metadata, $options);
    }

    /**
     * Update the details of given Restaurant
     * @param \Restaurant\Restaurant $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function UpdateRestaurant(\Restaurant\Restaurant $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/restaurant.RestaurantService/UpdateRestaurant',
        $argument,
        ['\Restaurant\Status', 'decode'],
        $metadata, $options);
    }

    /**
     * Delete Restaurant based on rest_id
     * @param \Restaurant\RestaurantId $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function DeleteRestaurant(\Restaurant\RestaurantId $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/restaurant.RestaurantService/DeleteRestaurant',
        $argument,
        ['\Restaurant\Status', 'decode'],
        $metadata, $options);
    }

    /**
     * Get all restaurants
     * @param \Restaurant\PBEmpty $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function GetAllRestaurants(\Restaurant\PBEmpty $argument,
      $metadata = [], $options = []) {
        return $this->_serverStreamRequest('/restaurant.RestaurantService/GetAllRestaurants',
        $argument,
        ['\Restaurant\Restaurant', 'decode'],
        $metadata, $options);
    }

}
