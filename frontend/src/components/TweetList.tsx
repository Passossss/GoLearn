
import { Tweet } from "@/services/TweetService";
import { Button } from "@/components/ui/button";
import { TrashIcon, InfoIcon } from "lucide-react";
import { formatDistanceToNow } from "date-fns";
import { ptBR } from "date-fns/locale";
import { Alert, AlertDescription } from "@/components/ui/alert";

interface TweetListProps {
  tweets: Tweet[];
  onDeleteTweet?: (id: string) => void;
  isLoading?: boolean;
}

export function TweetList({ tweets, onDeleteTweet, isLoading = false }: TweetListProps) {
  if (isLoading) {
    return (
      <div className="py-8 text-center text-muted-foreground">
        Carregando tweets...
      </div>
    );
  }

  const formatTweetTime = (dateString?: string) => {
    if (!dateString) return "";
    try {
      const date = new Date(dateString);
      return formatDistanceToNow(date, { addSuffix: true, locale: ptBR });
    } catch (e) {
      return "";
    }
  };

  return (
    <div className="space-y-4 mt-4">
      {tweets.length === 0 ? (
        <div className="py-8 text-center text-muted-foreground">
          Nenhum tweet encontrado. Seja o primeiro a postar!
        </div>
      ) : (
        <>
          <Alert className="bg-go-blue/10 border-go-blue/30 mb-4">
            <InfoIcon className="h-4 w-4 text-go-blue" />
            <AlertDescription className="text-sm">
              Esta é uma demo do TweetGO - Desenvolvida por Gustavo Passos!
            </AlertDescription>
          </Alert>
          
          {tweets.map((tweet) => (
            <div key={tweet.id} className="tweet-item p-4 border rounded-lg bg-card">
              <div className="flex justify-between">
                <div className="space-y-1">
                  <div className="text-sm text-muted-foreground">
                    {formatTweetTime(tweet.created_at)}
                  </div>
                  <div className="whitespace-pre-wrap break-words">
                    {tweet.description}
                  </div>
                </div>
                
                {onDeleteTweet && (
                  <Button
                    variant="ghost"
                    size="icon"
                    className="h-8 w-8 text-muted-foreground hover:text-destructive hover:bg-destructive/10"
                    onClick={() => onDeleteTweet(tweet.id)}
                    title="Deletar tweet"
                  >
                    <TrashIcon className="h-4 w-4" />
                  </Button>
                )}
              </div>
            </div>
          ))}
        </>
      )}
    </div>
  );
}
