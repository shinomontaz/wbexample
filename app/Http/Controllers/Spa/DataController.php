<?php
namespace App\Http\Controllers\Spa;
   
use Illuminate\Http\Request;
use App\Models\User;
use Illuminate\Http\JsonResponse;
use App\Http\Controllers\JsonController as JsonController;

use Illuminate\Support\Facades\Log;
use Illuminate\Support\Facades\Redis;

class DataController extends JsonController {
    
    public function __construct() {
    }
     /**
     * PauseFleet
     *
     * @return \Illuminate\Http\Response
     */
    public function GetTrucks(Request $request)
    {
        if (!Redis::exists('trucks')) {
            $jtrucks = json_encode('[]');
        } else {
            $jtrucks = Redis::get('trucks');
        }
        $trucks = json_decode($jtrucks, true);
        return $this->sendResponse($trucks);
    }
    
    public function GetPoints(Request $request)
    {
        if (!Redis::exists('points')) {
            $jpoints = json_encode('[]');
        } else {
            $jpoints = Redis::get('points');
        }
        $points = json_decode($jpoints, true);
        return $this->sendResponse($points);
    }
    
    public function GetPositions(Request $request)
    {
        if (!Redis::exists('positions')) {
            $jpositions = json_encode('[]');
        } else {
            $jpositions = Redis::get('positions');
        }
        $positions = json_decode($jpositions, true);
        return $this->sendResponse($positions);
    }

}