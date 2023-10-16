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

class CommandsController extends JsonController {
     /**
     * Login api
     *
     * @return \Illuminate\Http\Response
     */
    public function PauseFleet(Request $request): JsonResponse
    {
        $mqService = new RabbitMQService();
        $message = 'PauseFleet';
        $mqService->publish($message, 'wb', 'cmd');
        return $this->sendError('Unauthorised.', ['error'=>'Unauthorised']);
    }
    
    public function AddPoint(Request $request): JsonResponse
    {
        $mqService = new RabbitMQService();
        $message = 'add point';
        $mqService->publish($message, 'wb', 'cmd');
        return $this->sendError('Unauthorised.', ['error'=>'Unauthorised']);
    }
    
    public function GenerateFleet(Request $request): JsonResponse
    {
        $mqService = new RabbitMQService();
        $message = 'GenerateFleet';
        $mqService->publish($message, 'wb', 'cmd');
        return $this->sendError('Unauthorised.', ['error'=>'Unauthorised']);
    }    
}