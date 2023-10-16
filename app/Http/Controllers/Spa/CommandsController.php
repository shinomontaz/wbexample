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
        $message = [
            'type' => 1,
        ];
        return $this->_sendMessage($message);
    }
    
    public function AddPoint(Request $request)
    {
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
        $message = [
            'type' => 3,
        ];
        return $this->_sendMessage($message);
    }
    
    private function _sendMessage($message): JsonResponse {
        $message = json_encode($message);
        try {
            $this->mqService->publish($message, 'wb', 'cmd');
        } catch ( \Exception $e ) {
            return $this->sendError('Internal error.', ['error'=>$e->getMessage()], 500);
        }
        return $this->sendResponse('message sended.');
    }
}