<?php
namespace App\Http\Controllers\Spa;
   
use Illuminate\Http\Request;
use App\Models\User;
use Illuminate\Http\JsonResponse;
use App\Http\Controllers\JsonController as JsonController;

use Illuminate\Support\Facades\Log;
use Illuminate\Support\Facades\Cache;

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
        $jtrucks = Cache::get('trucks', json_encode([]));
        $trucks = json_decode($jtrucks, true);
        return $this->sendResponse($trucks);
    }
    
    public function GetPoints(Request $request)
    {
        $jpoints = Cache::get('points', json_encode([]));
        $points = json_decode($jpoints, true);
        return $this->sendResponse($points);
    }
    
    public function GetPositions(Request $request)
    {
        $jpositions = Cache::get('positions', json_encode([]));
        $positions = json_decode($jpositions, true);
        return $this->sendResponse($positions);
    }

}