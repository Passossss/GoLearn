
import { useState, useEffect } from "react";
import { Header } from "@/components/Header";
import { TweetBox } from "@/components/TweetBox";
import { TweetList } from "@/components/TweetList";
import { useTweetService, Tweet } from "@/services/TweetService";
import { useToast } from "@/components/ui/use-toast";

const Index = () => {
  const [tweets, setTweets] = useState<Tweet[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [isSubmitting, setIsSubmitting] = useState(false);
  const tweetService = useTweetService();
  const { toast } = useToast();

  const fetchTweets = async () => {
    setIsLoading(true);
    try {
      const fetchedTweets = await tweetService.getTweets();
      // Sort tweets by creation date, newest first
      const sortedTweets = fetchedTweets.sort((a, b) => {
        const dateA = a.created_at ? new Date(a.created_at) : new Date(0);
        const dateB = b.created_at ? new Date(b.created_at) : new Date(0);
        return dateB.getTime() - dateA.getTime();
      });
      setTweets(sortedTweets);
    } catch (error) {
      console.error("Error in fetchTweets:", error);
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    fetchTweets();
  }, []);

  const handleTweetSubmit = async (content: string) => {
    setIsSubmitting(true);
    try {
      await tweetService.postTweet(content);
      // Refresh the tweet list
      await fetchTweets();
    } catch (error) {
      console.error("Error in handleTweetSubmit:", error);
    } finally {
      setIsSubmitting(false);
    }
  };

  const handleDeleteTweet = async (id: string) => {
    const success = await tweetService.deleteTweet(id);
    if (success) {
      // Refresh the tweet list after deletion
      await fetchTweets();
    }
  };

  return (
    <div className="min-h-screen bg-background">
      <Header />
      
      <main className="container max-w-2xl py-6 px-4 md:px-6">
        <h1 className="text-2xl font-bold mb-6">Feed</h1>
        
        <div className="space-y-8">
          <div>
            <h2 className="text-lg font-medium mb-3">Compartilhe algo</h2>
            <TweetBox 
              onTweetSubmit={handleTweetSubmit} 
              isSubmitting={isSubmitting} 
            />
          </div>
          
          <div>
            <h2 className="text-lg font-medium mb-3">Timeline</h2>
            <TweetList 
              tweets={tweets} 
              onDeleteTweet={handleDeleteTweet} 
              isLoading={isLoading} 
            />
          </div>
        </div>
      </main>
    </div>
  );
};

export default Index;
