<?php
namespace App\Services;

use PhpAmqpLib\Connection\AMQPStreamConnection;
use PhpAmqpLib\Connection\AMQPSSLConnection;
use PhpAmqpLib\Message\AMQPMessage;

class RabbitMQService
{
    public function publish($message, $exchange, $rk)
    {
        $connection = new AMQPStreamConnection(env('MQ_HOST'), env('MQ_PORT'), env('MQ_USER'), env('MQ_PASS'), env('MQ_VHOST'));
        $channel = $connection->channel();
        $msg = new AMQPMessage($message);
        $channel->basic_publish($msg, $exchange, $rk);
        echo " [x] Sent $message to {$exchange} / {$rk}.\n";
        $channel->close();
        $connection->close();
    }
    public function consume($queue)
    {
        $connection = new AMQPStreamConnection(env('MQ_HOST'), env('MQ_PORT'), env('MQ_USER'), env('MQ_PASS'), env('MQ_VHOST'));
        $channel = $connection->channel();
        $callback = function ($msg) {
            echo ' [x] Received ', $msg->body, "\n";
        };
        $channel->basic_consume($queue, '', false, true, false, false, $callback);
        echo 'Waiting for new message on '.$queue. " \n";
        while ($channel->is_consuming()) {
            $channel->wait();
        }
        $channel->close();
        $connection->close();
    }
}