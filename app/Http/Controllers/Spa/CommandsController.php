<?php

/* 
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */

namespace App\Http\Controllers\Spa;
   
use Illuminate\Http\Request;
use App\Models\User;
use Illuminate\Support\Facades\Auth;
use Validator;
use Illuminate\Http\JsonResponse;
use \App\Services\RabbitMQService;
use App\Http\Controllers\JsonController as JsonController;

use Illuminate\Support\Facades\Log;
use Illuminate\Support\Facades\Redis;

class CommandsController extends JsonController {
    private $mqService;
    
    public function __construct() {
//        parent::__construct();
        $this->mqService = new RabbitMQService();
    }
     /**
     * PauseFleet
     *
     * @return \Illuminate\Http\Response
     */
    public function PauseFleet(Request $request)
    {
        //key here will be a "laravel_database_points"
//        Redis::set('points', json_encode([ 
//            ['id' => 211268, 'lat' => 55.547815000, 'long' => 37.550307000]
//        ]));
        $message = [
            'type' => 1,
        ];
        return $this->_sendMessage($message);
    }
    
    public function StartFleet(Request $request)
    {
        //{"type": 4}
        $message = [
            'type' => 4,
        ];
        return $this->_sendMessage($message);
    }
    
    public function AddPoint(Request $request)
    {
        //{location: [37.313919067382805, 55.879899669135995]}
//        Log::info( print_r($request->location, true) );
        $message = [
            'type' => 2,
            'long' => $request->location[0], // in SPA we use toLonLat function
            'lat' => $request->location[1],
        ];
        return $this->_sendMessage($message);
    }
    
    public function GenerateFleet(Request $request)
    {
        echo "sdfsdf sdfsdf";
        //{"num": 10, "type": 3, "area": {"Min": [37.43133544921874, 55.661286579672606], "Max": [37.56866455078125, 55.73867511243941]}}
        $message = [
            'type' => 3,
            'num' => $request->num,
            'area' => [
                'min' => [ $request->square[0], $request->square[1]],
                'max' => [ $request->square[2], $request->square[3]]
            ]
        ];
        return $this->_sendMessage($message);
    }
    
    private function _sendMessage($message): JsonResponse {
        $message = json_encode($message);
        print_r($message);
        die();
        try {
            $this->mqService->publish($message, 'wb', 'cmd');
        } catch ( \Exception $e ) {
            return $this->sendError('Internal error.', ['error'=>$e->getMessage()], 500);
        }
        return $this->sendResponse('message sended.');
    }
}